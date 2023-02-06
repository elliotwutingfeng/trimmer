# trimmer

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/elliotwutingfeng/trimmer)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotwutingfeng/trimmer?style=for-the-badge)](https://goreportcard.com/report/github.com/elliotwutingfeng/trimmer)

[![GitHub license](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)

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

fmt.Println(fastTrim("", cutset, TrimBoth))
fmt.Println(strings.Trim("", "@👍🏽新 "))

fmt.Println(fastTrim(" ", cutset, TrimBoth))
fmt.Println(strings.Trim(" ", "@👍🏽新 "))

fmt.Println(fastTrim("@b👍🏽新", cutset, TrimBoth))
fmt.Println(strings.Trim("@b👍🏽新", "@👍🏽新 "))

fmt.Println(fastTrim("@b👍🏽新", cutset, TrimLeft))
fmt.Println(strings.TrimLeft("@b👍🏽新", "@👍🏽新 "))

fmt.Println(fastTrim("@b👍🏽新", cutset, TrimRight))
fmt.Println(strings.TrimRight("@b👍🏽新", "@👍🏽新 "))

fmt.Println(fastTrim("@b👍新", cutset, TrimRight))
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

## Benchmarking

```sh
make bench
```