package cmd

import (
	"fmt"
	"os"

	. "github.com/iav0207/fcards/internal"
	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/game"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/model/mcard"

	"github.com/spf13/cobra"
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

	sample := game.RandomSampleOfMultiCardsFrom(cards)

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

func exitIfEmpty(cards []card.Card) {
	if len(cards) == 0 {
		Log.Println("Well, no game this time.")
		os.Exit(0)
	}
}

// Plays a round with given cards and returns those which were given wrong answers to.
func playRound(multicards []*mcard.MultiCard) []*mcard.MultiCard {
	wrongAnswered := make([]*mcard.MultiCard, 0)

	for _, mCard := range multicards {
		Log.Println("")
		response := UserResponse(mCard.Question)
		scored := game.Evaluate(*mCard, response)
		printGrade(scored)
		if scored.MissScore() > 0 {
			wrongAnswered = append(wrongAnswered, mCard)
		}
	}

	return wrongAnswered
}

func printGrade(sr game.Scored) {
	switch sr.MissScore() {
	case 0:
		Log.Println("✅")
	case 1, 2:
		Log.Println("🌼 Almost!", sr.Expected())
	default:
		Log.Println("🍅 Expected:", sr.Expected())
	}
	alternatives := sr.Alternatives()
	if len(alternatives) > 0 {
		Log.Println("Also valid:")
	}
	for _, alt := range alternatives {
		Log.Println(answerWithComment(alt))
	}
}

func answerWithComment(crd card.Card) string {
	if IsBlank(crd.Comment) {
		return crd.Answer
	}
	return fmt.Sprintf("%s (%s)", crd.Answer, crd.Comment)
}
