package game

import (
	. "github.com/iav0207/fcards/internal"
	"github.com/iav0207/fcards/internal/model"
)

func Evaluate(multicard model.MultiCard, response string) Scored {
	return Evaluator.Score(multicard, response)
}

type responseEvaluator interface {
	Score(multicard model.MultiCard, response string) scoredResponse
}

type Scored interface {
	MultiCard() model.MultiCard
	MissScore() int
	Expected() string
	Actual() string
	Alternatives() []model.Card
}

type evaluator struct{}

var Evaluator responseEvaluator = evaluator{}

// TODO score | assessment -> grade
func (ev evaluator) Score(multicard model.MultiCard, response string) scoredResponse {
	initMissScore := LevenshteinDistance(response, multicard.Cards[0].Answer)
	ret := scoredResponse{multicard, response, 0, initMissScore}
	for i, card := range multicard.Cards {
		score := LevenshteinDistance(response, card.Answer)
		if score < ret.missScore {
			ret.bestMatchIdx = i
			ret.missScore = score
		}
	}
	return ret
}

type scoredResponse struct {
	multicard    model.MultiCard
	response     string
	bestMatchIdx int
	missScore    int
}

func (sr scoredResponse) MultiCard() model.MultiCard {
	return sr.multicard
}

func (sr scoredResponse) MissScore() int {
	return sr.missScore
}

func (sr scoredResponse) Expected() string {
	return sr.multicard.Cards[sr.bestMatchIdx].Answer
}

func (sr scoredResponse) Actual() string {
	return sr.response
}

func (sr scoredResponse) Alternatives() []model.Card {
	alt := make([]model.Card, 0, len(sr.multicard.Cards)-1)
	for i, card := range sr.multicard.Cards {
		if i != sr.bestMatchIdx {
			alt = append(alt, card)
		}
	}
	return alt
}
