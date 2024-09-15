package game

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/model/mcard"
)

const testRandomSeed = 61409941995

func testRandom() rand.Rand { return *rand.New(rand.NewSource(testRandomSeed)) }

func TestSampler(t *testing.T) {
	t.Run("sample_size", func(t *testing.T) {
		want := 3
		sampler := &Sampler{
			sizeLimit: want,
			direc:     flags.Random,
			random:    testRandom(),
		}
		got := len(sampler.RandomSampleOfMultiCardsFrom(generateUniqueCards(7)))
		if got != want {
			t.Errorf("len(sample) = %d, want %d", got, want)
		}
	})

	t.Run("grouping", func(t *testing.T) {
		for _, tc := range []struct {
			name    string
			sampler *Sampler
			given   []card.Card
			want    []*mcard.MultiCard
		}{
			{
				name: "same_question",
				sampler: &Sampler{
					sizeLimit: 10,
					direc:     flags.Straight,
					random:    testRandom(),
				},
				given: generateCards(5, func(i int) *card.Card {
					c := cardNum(i)
					c.Question = "same question"
					return c
				}),
				want: []*mcard.MultiCard{
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
				},
			},
			{
				name: "deduplication",
				sampler: &Sampler{
					sizeLimit: 100,
					direc:     flags.Straight,
					random:    testRandom(),
				},
				given: []card.Card{
					*card.New("q1", "a1", "c1"),
					*card.New("q2", "a2", "c2"),
					*card.New("q3", "a3", "c3"),
					*card.New("q4", "a4", "c4"),
					*card.New("q5", "a5", "c5"),

					*card.New("q2", "a2", "c2"),
					*card.New("q3", "a3", "c3"),

					*card.New("q2", "a2", "c2"),
					*card.New("q2", "a2", ""),
					*card.New("q2", "a2*", "c2*"),
				},
				want: []*mcard.MultiCard{
					{Question: "q1", Cards: []card.Card{*card.New("q1", "a1", "c1")}},
					{Question: "q5", Cards: []card.Card{*card.New("q5", "a5", "c5")}},
					{Question: "q2", Cards: []card.Card{
						*card.New("q2", "a2", "c2"),
						*card.New("q2", "a2", ""),
						*card.New("q2", "a2*", "c2*"),
					}},
					{Question: "q4", Cards: []card.Card{*card.New("q4", "a4", "c4")}},
					{Question: "q3", Cards: []card.Card{*card.New("q3", "a3", "c3")}},
				},
			},
			{
				name: "group_items_are_unique",
				sampler: &Sampler{
					sizeLimit: 10,
					direc:     flags.Straight,
					random:    testRandom(),
				},
				given: generateCards(10, func(i int) *card.Card {
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
				}),
				want: []*mcard.MultiCard{
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
				},
			},
			{
				name: "sample_straight",
				sampler: &Sampler{
					sizeLimit: 10,
					direc:     flags.Straight,
					random:    *rand.New(rand.NewSource(testRandomSeed)),
				},
				given: generateUniqueCards(10),
				want: []*mcard.MultiCard{
					{Question: "q2", Cards: []card.Card{*cardNum(2)}},
					{Question: "q1", Cards: []card.Card{*cardNum(1)}},
					{Question: "q7", Cards: []card.Card{*cardNum(7)}},
					{Question: "q9", Cards: []card.Card{*cardNum(9)}},
					{Question: "q3", Cards: []card.Card{*cardNum(3)}},
					{Question: "q0", Cards: []card.Card{*cardNum(0)}},
					{Question: "q6", Cards: []card.Card{*cardNum(6)}},
					{Question: "q4", Cards: []card.Card{*cardNum(4)}},
					{Question: "q8", Cards: []card.Card{*cardNum(8)}},
					{Question: "q5", Cards: []card.Card{*cardNum(5)}},
				},
			},
			{
				name: "sample_inverse",
				sampler: &Sampler{
					sizeLimit: 10,
					direc:     flags.Inverse,
					random:    *rand.New(rand.NewSource(testRandomSeed)),
				},
				given: generateUniqueCards(10),
				want: []*mcard.MultiCard{
					{Question: "a2", Cards: []card.Card{*card.New("a2", "q2", "c2")}},
					{Question: "a1", Cards: []card.Card{*card.New("a1", "q1", "c1")}},
					{Question: "a7", Cards: []card.Card{*card.New("a7", "q7", "c7")}},
					{Question: "a9", Cards: []card.Card{*card.New("a9", "q9", "c9")}},
					{Question: "a3", Cards: []card.Card{*card.New("a3", "q3", "c3")}},
					{Question: "a0", Cards: []card.Card{*card.New("a0", "q0", "c0")}},
					{Question: "a6", Cards: []card.Card{*card.New("a6", "q6", "c6")}},
					{Question: "a4", Cards: []card.Card{*card.New("a4", "q4", "c4")}},
					{Question: "a8", Cards: []card.Card{*card.New("a8", "q8", "c8")}},
					{Question: "a5", Cards: []card.Card{*card.New("a5", "q5", "c5")}},
				},
			},
			{
				name: "sample_random",
				sampler: &Sampler{
					sizeLimit: 10,
					direc:     flags.Random,
					random:    *rand.New(rand.NewSource(testRandomSeed)),
				},
				given: generateUniqueCards(10),
				want: []*mcard.MultiCard{
					{Question: "a8", Cards: []card.Card{*card.New("a8", "q8", "c8")}},
					{Question: "q3", Cards: []card.Card{*cardNum(3)}},
					{Question: "a3", Cards: []card.Card{*card.New("a3", "q3", "c3")}},
					{Question: "q4", Cards: []card.Card{*cardNum(4)}},
					{Question: "a2", Cards: []card.Card{*card.New("a2", "q2", "c2")}},
					{Question: "a9", Cards: []card.Card{*card.New("a9", "q9", "c9")}},
					{Question: "q6", Cards: []card.Card{*cardNum(6)}},
					{Question: "q5", Cards: []card.Card{*cardNum(5)}},
					{Question: "q1", Cards: []card.Card{*cardNum(1)}},
					{Question: "q8", Cards: []card.Card{*cardNum(8)}},
				},
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				got := tc.sampler.RandomSampleOfMultiCardsFrom(tc.given)
				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Error("sampler.RandomSampleOfMultiCardsFrom(cards) result differs (-got, +want):\n" + diff)
				}
			})
		}
	})
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

func toString(n int) string { return fmt.Sprint(n) }
