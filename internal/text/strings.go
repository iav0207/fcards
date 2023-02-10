package text

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"strings"
)

func LevenshteinDistance(a, b string) int {
	opts := levenshtein.DefaultOptions
	return levenshtein.DistanceForStrings([]rune(a), []rune(b), opts)
}

func TabSeparated(values ...string) string {
	return strings.Join(values, "\t")
}

func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
