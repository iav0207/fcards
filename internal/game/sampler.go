package game

import (
	"math/rand"
	"sort"
	"time"

	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/config"
	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/model/mcard"
	"github.com/iav0207/fcards/internal/out"
)

type Sampler struct {
	sizeLimit int
	direc     flags.Direction
	random    rand.Rand
}

func NewSampler(direc flags.Direction) *Sampler {
	return &Sampler{
		sizeLimit: config.Get().GameDeckSize,
		direc:     direc,
		random:    *rand.New(rand.NewSource(seed())),
	}
}

func seed() int64 { return time.Now().UnixNano() }

func (s Sampler) RandomSampleOfMultiCardsFrom(cards []card.Card) []*mcard.MultiCard {
	var mcDirect index = createIndex(cards)
	var mcInverse index = createIndex(invert(cards))

	var limit int = min(len(mcDirect), len(mcInverse), s.sizeLimit)

	keysDirect := assignDirection(sortedKeysOf(mcDirect), flags.Straight)
	keysInverse := assignDirection(sortedKeysOf(mcInverse), flags.Inverse)

	var keyPool []directedQuestion
	if s.direc == flags.Random {
		keyPool = append(keyPool, keysDirect...)
		keyPool = append(keyPool, keysInverse...)
	} else if s.direc == flags.Straight {
		keyPool = append(keyPool, keysDirect...)
	} else if s.direc == flags.Inverse {
		keyPool = append(keyPool, keysInverse...)
	}

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

func sortedKeysOf(idx index) []string {
	keys := make([]string, len(idx))
	i := 0
	for key := range idx {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
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

func (s Sampler) shuffleQuestions(questions []directedQuestion) {
	s.random.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
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
