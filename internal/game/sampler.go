package game

import (
	"math/rand"
	"time"

	. "github.com/iav0207/fcards/internal"
	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/model"
)

type Sampler interface {
	RandomSampleOfMultiCardsFrom(cards []model.Card) []*model.MultiCard
}

type sampler struct {
	sizeLimit int
}

var sampleSizeLimit int = GetConfig().GameDeckSize
var SamplerService = sampler{sampleSizeLimit}

func RandomSampleOfMultiCardsFrom(cards []model.Card) []*model.MultiCard {
	return SamplerService.RandomSampleOfMultiCardsFrom(cards)
}

func (s sampler) RandomSampleOfMultiCardsFrom(cards []model.Card) []*model.MultiCard {
	var mcDirect index = createIndex(cards)
	var mcInverse index = createIndex(invert(cards))

	var limit int = min(len(mcDirect), len(mcInverse), s.sizeLimit)

	keysDirect := assignDirection(keysOf(mcDirect)[:limit], flags.Straight)
	keysInverse := assignDirection(keysOf(mcInverse)[:limit], flags.Inverse)

	var keyPool []directedQuestion = append(keysDirect, keysInverse...)
	shuffleQuestions(keyPool)
	keyPool = keyPool[:limit]
	sample := make([]*model.MultiCard, 0, limit)
	for _, key := range keyPool {
		var mCard model.MultiCard
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

type index = map[string]*model.MultiCard

func createIndex(cards []model.Card) index {
	return model.IndexMultiCards(model.ToMultiCards(cards))
}

func invert(cards []model.Card) []model.Card {
	inverted := make([]model.Card, 0, len(cards))
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
