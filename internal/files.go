package internal

import (
	"bufio"
	"fmt"
	"github.com/iav0207/fcards/internal/model"
	fpx "github.com/yargevad/filepathx"
	"os"
	s "strings"
)

var Home = os.Getenv("HOME")
var DefaultTsvFilesPattern = fmt.Sprintf("%s/.fcards/tsv/**/*.tsv", Home)
var AllTsvPaths = func() []string { return Glob(DefaultTsvFilesPattern) }

func ReadCardsFromPaths(paths []string) []model.Card {
	cards := make([]model.Card, 0)
	for _, filePath := range paths {
		Log.Printf("Reading cards from file %s\n", filePath)
		cards = append(cards, ReadCardsFromPath(filePath)...)
	}
	return cards
}

func ReadCardsFromPath(filePath string) []model.Card {
	cards := make([]model.Card, 0)
	for line := range LinesFrom(filePath) {
		if len(line) == 0 {
			continue
		}
		parsed, err := ParseCard(line)
		PanicIf(err)
		cards = append(cards, *parsed)
	}
	return cards
}

func ParseCard(line string) (*model.Card, error) {
	splitLine := s.Split(line, "\t")
	if len(splitLine) < 2 {
		return nil, fmt.Errorf(`Expected every non-empty line to be a tab-separated pair: question and answer.
		Got %s`, line)
	}
	question, answer := splitLine[0], splitLine[1]
	comment := ""
	if len(splitLine) > 2 {
		comment = splitLine[2]
	}
	return model.NewCard(question, answer, comment), nil
}

func LinesFrom(filePath string) chan string {
	file, err := os.Open(filePath)
	PanicIf(err)
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
	PanicIf(err)
	defer file.Close()
	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func AppendToFile(path string, lines ...string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	FatalIf(err, "Could not open the file")
	defer file.Close()
	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func Glob(glob string) []string {
	paths, err := fpx.Glob(glob)
	PanicIf(err)
	return paths
}
