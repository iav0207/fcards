package cmd

import (
	"fmt"
	. "github.com/iav0207/fcards/internal"
	"github.com/iav0207/fcards/internal/flags"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a quiz",
	Long:  `Start a quiz. Usage: fcards play file1.tsv`,
	Run:   runPlay,
}

var direc flags.Direction = flags.Straight

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().Var(&direc, direc.Name(), direc.HelpMsg())
}

func runPlay(cmd *cobra.Command, args []string) {
	Require(direc == flags.Random, "Only Random mode is supported at the moment")
	paths := argsOrAllTsvPaths(args)
	cards := ReadCardsFromPaths(paths)
	Log.Println("Read", len(cards), "cards in total.")
	exitIfEmpty(cards)

	sample := randomSampleOfMultiCardsFrom(cards, 20)

	Log.Println("Let's play!")
	reiterate := playRound(sample)
	playRound(reiterate)
}

func argsOrAllTsvPaths(args []string) []string {
	if len(args) > 0 {
		return args
	}
	return AllTsvPaths()
}

func exitIfEmpty(cards []Card) {
	if len(cards) == 0 {
		Log.Println("Well, no game this time.")
		os.Exit(0)
	}
}

// TODO refactor?
func randomSampleOfMultiCardsFrom(cards []Card, sampleSizeLimit int) []*MultiCard {
	mcDirect := IndexMultiCards(GroupCards(cards))
	mcInverse := IndexMultiCards(GroupCards(invert(cards)))

	limit := min(len(mcDirect), len(mcInverse), sampleSizeLimit)

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

// Plays a round with given cards and returns those which were given wrong answers to.
func playRound(multicards []*MultiCard) []*MultiCard {
	wrongAnswered := make([]*MultiCard, 0)

	for _, mCard := range multicards {
		Log.Println("")
		response := UserResponse(mCard.Question)
		// find card answer with min ldist
		scored := score(*mCard, response)
		printGrade(scored)
		if scored.missScore > 0 {
			wrongAnswered = append(wrongAnswered, mCard)
		}
	}

	return wrongAnswered
}

// TODO score | assessment -> grade
func score(mCard MultiCard, response string) scoredResponse {
	initMissScore := LevenshteinDistance(response, mCard.Cards[0].Answer)
	ret := scoredResponse{mCard, response, 0, initMissScore}
	for i, card := range mCard.Cards {
		score := LevenshteinDistance(response, card.Answer)
		if score < ret.missScore {
			ret.bestMatchIdx = i
			ret.missScore = score
		}
	}
	return ret
}

type scoredResponse struct {
	multicard    MultiCard
	response     string
	bestMatchIdx int
	missScore    int
}

func (sr scoredResponse) expected() string {
	return sr.multicard.Cards[sr.bestMatchIdx].Answer
}

func (sr scoredResponse) alternatives() []Card {
	alt := make([]Card, 0, len(sr.multicard.Cards)-1)
	for i, card := range sr.multicard.Cards {
		if i != sr.bestMatchIdx {
			alt = append(alt, card)
		}
	}
	return alt
}

func printGrade(sr scoredResponse) {
	switch sr.missScore {
	case 0:
		Log.Println("✅")
	case 1, 2:
		Log.Println("🌼 Almost!", sr.expected())
	default:
		Log.Println("🍅 Expected:", sr.expected())
	}
	alternatives := sr.alternatives()
	if len(alternatives) > 0 {
		Log.Println("Also valid:")
	}
	for _, alt := range alternatives {
		Log.Println(answerWithComment(alt))
	}
}

func answerWithComment(card Card) string {
	if IsBlank(card.Comment) {
		return card.Answer
	}
	return fmt.Sprintf("%s (%s)", card.Answer, card.Comment)
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
