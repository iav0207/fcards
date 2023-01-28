package internal

import "strings"

type Card struct {
	Question, Answer string
}

func NewCard(q, a string) *Card {
	return &Card{q, a}
}

func (card *Card) Invert() {
	card.Question, card.Answer = card.Answer, card.Question
}

func (card Card) String() string {
	return strings.Join([]string{card.Question, card.Answer}, "\t")
}
