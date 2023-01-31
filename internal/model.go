package internal

import (
	"strings"
)

type Card struct {
	Question, Answer, Comment string
}

// Represents a card with one or more correct answers.
type MultiCard struct {
	Question string
	Cards    []Card
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

func NewMultiCard(q string, cards []Card) *MultiCard {
	set := make(map[string]*Card)
	for _, card := range cards {
		set[card.String()] = &card
	}
	uniqueCards := make([]Card, len(set))
	i := 0
	for _, card := range set {
		uniqueCards[i] = *card
		i++
	}
	return &MultiCard{q, cards}
}

func IndexMultiCards(multicards []*MultiCard) map[string]*MultiCard {
	index := make(map[string]*MultiCard)
	for _, mCard := range multicards {
		index[mCard.Question] = mCard
	}
	return index
}

func GroupCards0(cards []Card) map[string][]Card {
	multimap := make(map[string][]Card)
	for _, card := range cards {
		multimap[card.Question] = append(multimap[card.Question], card)
	}
	return multimap
}

func GroupCards(cards []Card) []*MultiCard {
	multimap := GroupCards0(cards)
	multicards := make([]*MultiCard, 0, len(multimap))
	for q, cards := range multimap {
		multicards = append(multicards, NewMultiCard(q, cards))
	}
	return multicards
}
