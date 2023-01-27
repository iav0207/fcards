package cmd

import (
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
	Run:   run,
}

var direc flags.Direction = flags.Straight

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().Var(&direc, direc.Name(), direc.HelpMsg())
}

func run(cmd *cobra.Command, args []string) {
	paths := argsOrAllTsvPaths(args)
	cards := ReadCardsFromPaths(paths)
	Log.Println("Read", len(cards), "cards in total.")
	exitIfEmpty(cards)

	sample := randomSampleOf(cards, 20)
	applyDirectionFlag(sample)

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

func randomSampleOf(cards []Card, sizeLimit int) []Card {
	shuffle(cards)
	sample := cards[:min(len(cards), sizeLimit)]
	if len(sample) != len(cards) {
		Log.Println("Took a random sample of", len(sample), "cards.")
	}
	return sample
}

// Plays a round with given cards and returns those which were given wrong answers to.
func playRound(cards []Card) []Card {
	wrongAnswered := make([]Card, 0)

	for _, card := range cards {
		Log.Println(card.Question)
		missScore := LevenshteinDistance(ReadLine(), card.Answer)
		printResponse(missScore, card.Answer)
		if missScore > 0 {
			wrongAnswered = append(wrongAnswered, card)
		}
	}

	return wrongAnswered
}

func printResponse(missScore int, expected string) {
	switch missScore {
	case 0:
		Log.Println("✅")
	case 1, 2:
		Log.Println("🌼 Almost!", expected)
	default:
		Log.Println("🍅 Expected:", expected)
	}
}

func shuffle(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func applyDirectionFlag(cards []Card) {
	for i := 0; i < len(cards); i++ {
		if shouldInvert() {
			cards[i].Invert()
		}
	}
}

func shouldInvert() bool {
	return direc == flags.Inverse || (direc == flags.Random && randomBool())
}

var randomBool = func() bool { return rand.Intn(2) == 0 }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
