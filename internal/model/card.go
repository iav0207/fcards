package model

import "strings"

type Card struct {
	Question, Answer, Comment string
}

func NewCard(question, answer, comment string) *Card {
	return &Card{question, answer, comment}
}

func (card Card) Copy() *Card {
	return NewCard(card.Question, card.Answer, card.Comment)
}

func (card *Card) Invert() {
	card.Question, card.Answer = card.Answer, card.Question
}

func (card Card) String() string {
	tabSeparated := strings.Join([]string{card.Question, card.Answer, card.Comment}, "\t")
	return strings.TrimRight(tabSeparated, " \t")
}
