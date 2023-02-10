package card

import (
	"fmt"
	"strings"
)

type Card struct {
	Question, Answer, Comment string
}

func New(question, answer, comment string) *Card {
	return &Card{question, answer, comment}
}

func Parse(line string) (*Card, error) {
	splitLine := strings.Split(line, "\t")
	if len(splitLine) < 2 {
		return nil, fmt.Errorf(`Expected every non-empty line to be a tab-separated pair: question and answer. Got: %s`, line)
	}
	question, answer := splitLine[0], splitLine[1]
	comment := ""
	if len(splitLine) > 2 {
		comment = splitLine[2]
	}
	return New(question, answer, comment), nil
}

func (card Card) Copy() *Card {
	return New(card.Question, card.Answer, card.Comment)
}

func (card *Card) Invert() {
	card.Question, card.Answer = card.Answer, card.Question
}

func (card Card) String() string {
	tabSeparated := strings.Join([]string{card.Question, card.Answer, card.Comment}, "\t")
	return strings.TrimRight(tabSeparated, " \t")
}
