package data

import (
	"bufio"
	"fmt"
	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/out"
	fpx "github.com/yargevad/filepathx"
	"os"
)

var Home = os.Getenv("HOME")
var DefaultTsvFilesPattern = fmt.Sprintf("%s/.fcards/tsv/**/*.tsv", Home)
var AllTsvPaths = func() []string { return Glob(DefaultTsvFilesPattern) }

func ReadCardsFromPaths(paths []string) []card.Card {
	cards := make([]card.Card, 0)
	for _, filePath := range paths {
		out.Log.Printf("Reading cards from file %s\n", filePath)
		cards = append(cards, ReadCardsFromPath(filePath)...)
	}
	return cards
}

func ReadCardsFromPath(filePath string) []card.Card {
	cards := make([]card.Card, 0)
	for line := range LinesFrom(filePath) {
		if len(line) == 0 {
			continue
		}
		parsed, err := card.Parse(line)
		check.FatalIf(err, "Failed to load the cards.")
		cards = append(cards, *parsed)
	}
	return cards
}

func LinesFrom(filePath string) chan string {
	file, err := os.Open(filePath)
	check.PanicIf(err)
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)
	c := make(chan string)
	go func() {
		defer file.Close()
		defer close(c)
		for sc.Scan() {
			c <- sc.Text()
		}
	}()
	return c
}

func OverwriteFileWithLines(path string, lines []string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	check.PanicIf(err)
	defer file.Close()
	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func AppendToFile(path string, lines ...string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check.FatalIf(err, "Could not open the file")
	defer file.Close()
	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func Glob(glob string) []string {
	paths, err := fpx.Glob(glob)
	check.PanicIf(err)
	return paths
}
