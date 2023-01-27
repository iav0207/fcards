package internal

import (
	"fmt"
	"os"
	fp "path/filepath"
	s "strings"
)

var Home = os.Getenv("HOME")

func ReadCardsFromPaths(paths []string) []Card {
	cards := make([]Card, 0)
	for _, filePath := range paths {
		Log.Printf("Reading cards from file %s\n", filePath)
		cards = append(cards, ReadCardsFromPath(filePath)...)
	}
	return cards
}

func ReadCardsFromPath(filePath string) []Card {
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
	return Glob(fmt.Sprintf("%s/.fcards/tsv/*.tsv", Home))
}

func Glob(glob string) []string {
	paths, err := fp.Glob(glob)
	check(err)
	return paths
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
