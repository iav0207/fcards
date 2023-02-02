package model

// Represents a card with one or more correct answers.
type MultiCard struct {
	Question string
	Cards    []Card
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

func ToMultiCards(cards []Card) []*MultiCard {
	multimap := GroupByQuestion(cards)
	multicards := make([]*MultiCard, 0, len(multimap))
	for q, cards := range multimap {
		multicards = append(multicards, NewMultiCard(q, deduplicate(cards)))
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
	set := make(map[Card]struct{})
	member := struct{}{}
	for _, card := range cards {
		set[card] = member
	}
	slice := make([]Card, len(set))
	i := 0
	for card := range set {
		slice[i] = card
		i++
	}
	return slice
}
