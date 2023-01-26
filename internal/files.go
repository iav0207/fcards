package internal

import (
	"os"
	fp "path/filepath"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadCards(filePath string) []Card {
	cards := make([]Card, 0)
	for _, line := range ReadLines(filePath) {
		if len(line) == 0 {
			continue
		}
		splitLine := s.Split(line, "\t")
		if len(splitLine) < 2 {
			panic("Expected every non-empty line to be a tab-separated pair: question and answer")
		}
		cards = append(cards, *NewCard(splitLine[0], splitLine[1]))
	}
	return cards
}

func ReadLines(filePath string) []string {
	return s.Split(ReadText(filePath), "\n")
}

func ReadText(filePath string) string {
	data, err := os.ReadFile(filePath)
	check(err)
	return string(data)
}

func AllTsvPaths() []string {
	pattern := os.Getenv("HOME") + "/.fcards/tsv/*.tsv"
	paths, err := fp.Glob(pattern)
	check(err)
	return paths
}
