package mcard

import "github.com/iav0207/fcards/internal/model/card"

// Represents a card with one or more correct answers.
type MultiCard struct {
	Question string
	Cards    []card.Card
}

func New(q string, cards []card.Card) *MultiCard {
	return &MultiCard{q, deduplicate(cards)}
}

func IndexMultiCards(multicards []*MultiCard) map[string]*MultiCard {
	index := make(map[string]*MultiCard)
	for _, mCard := range multicards {
		index[mCard.Question] = mCard
	}
	return index
}

func ToMultiCards(cards []card.Card) []*MultiCard {
	multimap := GroupByQuestion(cards)
	multicards := make([]*MultiCard, 0, len(multimap))
	for q, qCards := range multimap {
		multicards = append(multicards, New(q, qCards))
	}
	return multicards
}

func GroupByQuestion(cards []card.Card) map[string][]card.Card {
	multimap := make(map[string][]card.Card)
	for _, c := range cards {
		multimap[c.Question] = append(multimap[c.Question], c)
	}
	return multimap
}

func deduplicate(cards []card.Card) []card.Card {
	type member = struct{}
	set := make(map[card.Card]member)
	result := make([]card.Card, 0)
	for _, c := range cards {
		_, seen := set[c]
		if !seen {
			result = append(result, c)
			set[c] = member{}
		}
	}
	return result
}
