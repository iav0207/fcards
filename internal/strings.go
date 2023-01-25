package internal

import (
	"runtime"
	"strings"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func LevenshteinDistance(a, b string) int {
	ar := []rune(normalize(a))
	br := []rune(normalize(b))
	opts := levenshtein.DefaultOptions
	return levenshtein.DistanceForStrings(ar, br, opts)
}

func Compare(a, b string) int {
	return strings.Compare(normalize(a), normalize(b))
}

func normalize(s string) string {
	return strings.TrimRight(s, lineBreak())
}

func lineBreak() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}
