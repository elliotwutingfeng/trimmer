// Package trimmer is an alternative implementation of the standard Go
// strings.Trim, strings.TrimLeft, and strings.TrimRight functions.
package trimmer

import (
	"unicode/utf8"

	"github.com/karlseguin/intset"
)

// TrimMode specifies which parts of string to trim for FastTrim()
type TrimMode int

const (
	// TrimBoth trims both ends of the string
	TrimBoth TrimMode = iota
	// TrimLeft trims the left end of the string
	TrimLeft
	// TrimRight trims the right end of the string
	TrimRight
)

// FastTrim works like strings.Trim but uses cutset of type *intset.Rune
func FastTrim(s string, cutset *intset.Rune, mode TrimMode) string {
	var startIdx, endIdx int
	if mode != TrimRight {
		// Trim left-hand side
		var trimCharsExist bool
		var broken bool
		for idx, c := range s {
			startIdx = idx
			if !cutset.Exists(c) {
				broken = true
				break
			}
			trimCharsExist = true
		}
		if trimCharsExist && !broken {
			// Return empty string if every character in s exists in cutset
			return ""
		}
	}
	if mode != TrimLeft {
		// Trim right-hand side
		var trimCharsExist bool
		var broken bool
		for i := len(s); i > 0; {
			endIdx = i
			r, size := utf8.DecodeLastRuneInString(s[0:i])
			i -= size
			if !cutset.Exists(r) {
				broken = true
				break
			}
			trimCharsExist = true
		}
		if trimCharsExist && !broken {
			// Return empty string if every character in s exists in cutset
			return ""
		}
	} else {
		endIdx = len(s)
	}
	return s[startIdx:endIdx]
}

// MakeRuneSet converts a string to a set of unique runes
func MakeRuneSet(s string) (iset *intset.Rune) {
	var biggestRune rune
	for idx, r := range s {
		if idx == 0 || r > biggestRune {
			biggestRune = r
		}
	}
	// optimal target capacity
	iset = intset.NewRune(biggestRune)
	for _, r := range s {
		iset.Set(r)
	}
	return
}
