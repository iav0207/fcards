package game

import (
	"math/rand"
	"testing"

	"github.com/go-test/deep"

	"fmt"

	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/model/mcard"
)

const testRandomSeed = 61409941995

func testRandom() rand.Rand { return *rand.New(rand.NewSource(testRandomSeed)) }

func TestSampleShouldHaveExpectedSize(t *testing.T) {
	expected := 3
	sampler := createTestSampler(expected)
	actual := len(sampler.RandomSampleOfMultiCardsFrom(generateUniqueCards(7)))
	if actual != expected {
		t.Errorf("len(sample) = %d, expected %d", actual, expected)
	}
}

func TestSampleGroupingDuplicatesCollapsed(t *testing.T) {
	sampler := createTestSampler(20)
	expectedCount := 5
	cards := generateUniqueCards(expectedCount)
	cards = append(cards, cards...)
	cards = append(cards, cards[0])
	sample := sampler.RandomSampleOfMultiCardsFrom(cards)
	if len(sample) != expectedCount {
		t.Errorf("Expected %d groups only, got %d", expectedCount, len(sample))
	}
}

func TestGroupingForOneQuestion(t *testing.T) {
	sampler := &Sampler{
		sizeLimit: 10,
		direc:     flags.Straight,
		random:    testRandom(),
	}
	cards := generateCards(5, func(i int) *card.Card {
		c := cardNum(i)
		c.Question = "same question"
		return c
	})
	sample := sampler.RandomSampleOfMultiCardsFrom(cards)
	if len(sample) != 1 { // this test depends on the random seed
		t.Errorf("Expected just one group, got %d", len(sample))
	}
	expected := []*mcard.MultiCard{
		{
			Question: "same question",
			Cards: []card.Card{
				*card.New("same question", "a0", "c0"),
				*card.New("same question", "a1", "c1"),
				*card.New("same question", "a2", "c2"),
				*card.New("same question", "a3", "c3"),
				*card.New("same question", "a4", "c4"),
			},
		},
	}
	if diff := deep.Equal(sample, expected); diff != nil {
		t.Error(diff)
	}
}

func TestGroupItemsAreUnique(t *testing.T) {
	sampler := &Sampler{
		sizeLimit: 10,
		direc:     flags.Straight,
		random:    testRandom(),
	}
	cards := generateCards(10, func(i int) *card.Card {
		crComm := func() string {
			if i%3 == 0 {
				return "c" + toString(i%2)
			} else {
				return ""
			}
		}
		return &card.Card{
			Question: "q" + toString(i%2),
			Answer:   "a" + toString(i%3),
			Comment:  crComm(),
		}
	})
	sample := sampler.RandomSampleOfMultiCardsFrom(cards)
	expected := []*mcard.MultiCard{
		{
			Question: "q0",
			Cards: []card.Card{
				*card.New("q0", "a0", "c0"),
				*card.New("q0", "a2", ""),
				*card.New("q0", "a1", ""),
			},
		},
		{
			Question: "q1",
			Cards: []card.Card{
				*card.New("q1", "a1", ""),
				*card.New("q1", "a0", "c1"),
				*card.New("q1", "a2", ""),
			},
		},
	}
	if diff := deep.Equal(sample, expected); diff != nil {
		t.Error(diff)
	}
}

func TestStraightDirectionSampling(t *testing.T) {
	sampler := &Sampler{
		sizeLimit: 10,
		direc:     flags.Straight,
		random:    *rand.New(rand.NewSource(testRandomSeed)),
	}
	cards := generateCards(10, func(i int) *card.Card {
		return &card.Card{
			Question: "q" + toString(i),
			Answer:   "a" + toString(i),
			Comment:  "",
		}
	})
	sample := sampler.RandomSampleOfMultiCardsFrom(cards)
	expected := []*mcard.MultiCard{
		{Question: "q2", Cards: []card.Card{*card.New("q2", "a2", "")}},
		{Question: "q1", Cards: []card.Card{*card.New("q1", "a1", "")}},
		{Question: "q7", Cards: []card.Card{*card.New("q7", "a7", "")}},
		{Question: "q9", Cards: []card.Card{*card.New("q9", "a9", "")}},
		{Question: "q3", Cards: []card.Card{*card.New("q3", "a3", "")}},
		{Question: "q0", Cards: []card.Card{*card.New("q0", "a0", "")}},
		{Question: "q6", Cards: []card.Card{*card.New("q6", "a6", "")}},
		{Question: "q4", Cards: []card.Card{*card.New("q4", "a4", "")}},
		{Question: "q8", Cards: []card.Card{*card.New("q8", "a8", "")}},
		{Question: "q5", Cards: []card.Card{*card.New("q5", "a5", "")}},
	}
	if diff := deep.Equal(sample, expected); diff != nil {
		t.Error(diff)
	}
}

func createTestSampler(sizeLimit int) *Sampler {
	return &Sampler{
		sizeLimit: sizeLimit,
		direc:     flags.Random,
		random:    testRandom(),
	}
}

func generateUniqueCards(count int) []card.Card {
	var generate = func(i int) *card.Card { return cardNum(i) }
	return generateCards(count, generate)
}

func generateCards(count int, generate func(int) *card.Card) []card.Card {
	cards := make([]card.Card, count)
	for i := 0; i < count; i++ {
		cards[i] = *generate(i)
	}
	return cards
}

func cardNum(num int) *card.Card {
	id := toString(num)
	return card.New("q"+id, "a"+id, "c"+id)
}

func toString(n int) string {
	return fmt.Sprint(n)
}
