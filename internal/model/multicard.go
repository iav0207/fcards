package model

// Represents a card with one or more correct answers.
type MultiCard struct {
	Question string
	Cards    []Card
}

func NewMultiCard(q string, cards []Card) *MultiCard {
	return &MultiCard{q, deduplicate(cards)}
}

func IndexMultiCards(multicards []*MultiCard) map[string]*MultiCard {
	index := make(map[string]*MultiCard)
	for _, mCard := range multicards {
		index[mCard.Question] = mCard
	}
	return index
}

func ToMultiCards(cards []Card) []*MultiCard {
	multimap := GroupByQuestion(cards)
	multicards := make([]*MultiCard, 0, len(multimap))
	for q, qCards := range multimap {
		multicards = append(multicards, NewMultiCard(q, qCards))
	}
	return multicards
}

func GroupByQuestion(cards []Card) map[string][]Card {
	multimap := make(map[string][]Card)
	for _, card := range cards {
		multimap[card.Question] = append(multimap[card.Question], card)
	}
	return multimap
}

func deduplicate(cards []Card) []Card {
	type member = struct{}
	set := make(map[Card]member)
	result := make([]Card, 0)
	for _, card := range cards {
		_, seen := set[card]
		if !seen {
			result = append(result, card)
			set[card] = member{}
		}
	}
	return result
}
