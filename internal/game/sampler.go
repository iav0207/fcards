package game

import (
	"math/rand"
	"time"

	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/config"
	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/model/mcard"
	"github.com/iav0207/fcards/internal/out"
)

type Sampler interface {
	RandomSampleOfMultiCardsFrom(cards []card.Card) []*mcard.MultiCard
}

type sampler struct {
	sizeLimit  int
	randomSeed int64
}

var SamplerService = sampler{
	sizeLimit:  config.Get().GameDeckSize,
	randomSeed: time.Now().UnixNano(),
}

func RandomSampleOfMultiCardsFrom(cards []card.Card) []*mcard.MultiCard {
	return SamplerService.RandomSampleOfMultiCardsFrom(cards)
}

func (s sampler) RandomSampleOfMultiCardsFrom(cards []card.Card) []*mcard.MultiCard {
	var mcDirect index = createIndex(cards)
	var mcInverse index = createIndex(invert(cards))

	var limit int = min(len(mcDirect), len(mcInverse), s.sizeLimit)

	keysDirect := assignDirection(keysOf(mcDirect)[:limit], flags.Straight)
	keysInverse := assignDirection(keysOf(mcInverse)[:limit], flags.Inverse)

	var keyPool []directedQuestion = append(keysDirect, keysInverse...)
	s.shuffleQuestions(keyPool)
	keyPool = keyPool[:limit]
	sample := make([]*mcard.MultiCard, 0, limit)
	for _, key := range keyPool {
		var mCard mcard.MultiCard
		if key.direc == flags.Straight {
			mCard = *mcDirect[key.question]
		} else {
			mCard = *mcInverse[key.question]
		}
		sample = append(sample, &mCard)
	}
	if len(sample) != len(cards) {
		out.Log.Println("Took a random sample of", len(sample), "cards")
	}
	return sample
}

type index = map[string]*mcard.MultiCard

func createIndex(cards []card.Card) index {
	return mcard.IndexMultiCards(mcard.ToMultiCards(cards))
}

func invert(cards []card.Card) []card.Card {
	inverted := make([]card.Card, 0, len(cards))
	for _, c := range cards {
		c.Invert()
		inverted = append(inverted, c)
	}
	return inverted
}

type directedQuestion struct {
	question string
	direc    flags.Direction
}

func keysOf(idx index) []string {
	keys := make([]string, len(idx))
	i := 0
	for key := range idx {
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

func (s sampler) shuffleQuestions(questions []directedQuestion) {
	rand.Seed(s.randomSeed)
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
}

func min(items ...int) int {
	check.Assert(len(items) > 0)
	ret := items[0]
	for i := 1; i < len(items); i++ {
		if items[i] < ret {
			ret = items[i]
		}
	}
	return ret
}
