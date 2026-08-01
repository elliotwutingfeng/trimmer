# trimmer

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/elliotwutingfeng/trimmer)
[![Coveralls](https://img.shields.io/coverallsCoverage/github/elliotwutingfeng/trimmer?logo=coveralls&style=for-the-badge)](https://coveralls.io/github/elliotwutingfeng/trimmer?branch=main)

[![License](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)

## Summary

**trimmer** is an alternative implementation of the standard Go [strings.Trim](https://pkg.go.dev/strings#Trim), [strings.TrimLeft](https://pkg.go.dev/strings#TrimLeft), and [strings.TrimRight](https://pkg.go.dev/strings#TrimRight) functions.

Runes to be trimmed away are stored in a [bitset](https://github.com/karlseguin/intset) for quick retrieval.

On average, **trimmer** executes at least *twice* as fast as [strings.Trim](https://pkg.go.dev/strings#Trim).

Spot any bugs? Report them [here](https://github.com/elliotwutingfeng/trimmer/issues).

## Installation

```sh
go get github.com/elliotwutingfeng/trimmer
```

## Example

```go
const charsToTrim string = "@👍🏽新 "
var cutset *intset.Rune = MakeRuneSet(charsToTrim)

fmt.Println(FastTrim("", cutset, TrimBoth))
fmt.Println(strings.Trim("", "@👍🏽新 "))

fmt.Println(FastTrim(" ", cutset, TrimBoth))
fmt.Println(strings.Trim(" ", "@👍🏽新 "))

fmt.Println(FastTrim("@b👍🏽新", cutset, TrimBoth))
fmt.Println(strings.Trim("@b👍🏽新", "@👍🏽新 "))

fmt.Println(FastTrim("@b👍🏽新", cutset, TrimLeft))
fmt.Println(strings.TrimLeft("@b👍🏽新", "@👍🏽新 "))

fmt.Println(FastTrim("@b👍🏽新", cutset, TrimRight))
fmt.Println(strings.TrimRight("@b👍🏽新", "@👍🏽新 "))

fmt.Println(FastTrim("@b👍新", cutset, TrimRight))
fmt.Println(strings.TrimRight("@b👍新", "@👍🏽新 "))
//Output:
//
//
//b
//b
//b👍🏽新
//b👍🏽新
//@b
//@b
//@b
//@b
```

## Testing

```sh
make tests

# Alternatively, run tests without race detection
# Useful for systems that do not support the -race flag like windows/386
# See https://tip.golang.org/src/cmd/dist/test.go
make tests_without_race
```

## Benchmarks

```sh
make bench
```

### Results

The following chart indicates total time taken to trim characters `@👍🏽新` from strings in <https://github.com/minimaxir/big-list-of-naughty-strings>. The strings have been modified to include prefix and/or suffix permutations with replacement of up to length 4 from the set `@👍🏽新`.

```text
CPU: AMD Ryzen 7 5800X
Time in milliseconds (ms) | Lower is better

  MakeRuneSet ▏ .08

     FastTrim ▏ 252 🟦🟦🟦🟦🟦🟦🟦🟦🟦

 FastTrimLeft ▏ 113 🟦🟦🟦🟦

FastTrimRight ▏ 167 🟦🟦🟦🟦🟦🟦

         Trim ▏ 680 🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥

     TrimLeft ▏ 342 🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥

    TrimRight ▏ 382 🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥
```

## Limitations

The cutset (characters to be trimmed) must be initialized as `*intset.Rune` first. This incurs a small time cost.

```go
const charsToTrim string = "@👍🏽新 "
var cutset *intset.Rune = MakeRuneSet(charsToTrim) // this incurs a small time cost
```

Use the Go Standard library if you only need to process a small amount of strings (10 strings or less) per given cutset. However, for larger amounts of strings per given cutset, **trimmer** will greatly outperform the standard library.
