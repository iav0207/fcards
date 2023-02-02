package game

import (
	"math/rand"
	"time"

	. "github.com/iav0207/fcards/internal"
	"github.com/iav0207/fcards/internal/flags"
)

type Sampler interface {
	RandomSampleOfMultiCardsFrom(cards []Card) []*MultiCard
}

type sampler struct {
	sizeLimit int
}

var SamplerService = sampler{20} // TODO config?

func RandomSampleOfMultiCardsFrom(cards []Card) []*MultiCard {
	return SamplerService.RandomSampleOfMultiCardsFrom(cards)
}

func (s sampler) RandomSampleOfMultiCardsFrom(cards []Card) []*MultiCard {
	mcDirect := IndexMultiCards(ToMultiCards(cards))
	mcInverse := IndexMultiCards(ToMultiCards(invert(cards)))

	limit := min(len(mcDirect), len(mcInverse), s.sizeLimit)

	keysDirect := assignDirection(keysOf(mcDirect)[:limit], flags.Straight)
	keysInverse := assignDirection(keysOf(mcInverse)[:limit], flags.Inverse)

	keyPool := append(keysDirect, keysInverse...)
	shuffleQuestions(keyPool)
	keyPool = keyPool[:limit]
	sample := make([]*MultiCard, 0, limit)
	for _, key := range keyPool {
		var mCard MultiCard
		if key.direc == flags.Straight {
			mCard = *mcDirect[key.question]
		} else {
			mCard = *mcInverse[key.question]
		}
		sample = append(sample, &mCard)
	}
	if len(sample) != len(cards) {
		Log.Println("Took a random sample of", len(sample), "cards")
	}
	return sample
}

func invert(cards []Card) []Card {
	inverted := make([]Card, 0, len(cards))
	for _, card := range cards {
		card.Invert()
		inverted = append(inverted, card)
	}
	return inverted
}

type directedQuestion struct {
	question string
	direc    flags.Direction
}

func keysOf(m map[string]*MultiCard) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}

func assignDirection(questions []string, direc flags.Direction) []directedQuestion {
	ret := make([]directedQuestion, len(questions))
	i := 0
	for _, q := range questions {
		ret[i] = directedQuestion{q, direc}
		i++
	}
	return ret
}

func shuffleQuestions(questions []directedQuestion) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
}

func min(items ...int) int {
	Assert(len(items) > 0)
	ret := items[0]
	for i := 1; i < len(items); i++ {
		if items[i] < ret {
			ret = items[i]
		}
	}
	return ret
}
