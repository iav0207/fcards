package cmd

import (
	"fmt"
	"os"

	"github.com/iav0207/fcards/internal/data"
	"github.com/iav0207/fcards/internal/flags"
	"github.com/iav0207/fcards/internal/game"
	"github.com/iav0207/fcards/internal/in"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/model/mcard"
	"github.com/iav0207/fcards/internal/out"
	"github.com/iav0207/fcards/internal/text"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a quiz",
	Long:  `Start a quiz. Usage: fcards play file1.tsv`,
	Run:   runPlay,
}

var direc flags.Direction = flags.Random

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().Var(&direc, direc.Name(), direc.HelpMsg())
}

func runPlay(cmd *cobra.Command, args []string) {
	paths := argsOrAllTsvPaths(args)
	cards := data.ReadCardsFromPaths(paths)
	out.Log.Println("Read", len(cards), "cards in total.")
	exitIfEmpty(cards)

	sample := game.NewSampler(direc).RandomSampleOfMultiCardsFrom(cards)

	out.Log.Println("Let's play!")
	reiterate := playRound(sample)
	playRound(reiterate)
}

func argsOrAllTsvPaths(args []string) []string {
	if len(args) > 0 {
		return args
	}
	return data.AllTsvPaths()
}

func exitIfEmpty(cards []card.Card) {
	if len(cards) == 0 {
		out.Log.Println("Well, no game this time.")
		os.Exit(0)
	}
}

// Plays a round with given cards and returns those which were given wrong answers to.
func playRound(multicards []*mcard.MultiCard) []*mcard.MultiCard {
	wrongAnswered := make([]*mcard.MultiCard, 0)

	for _, mCard := range multicards {
		out.Log.Println("")
		response := in.UserResponse(mCard.Question)
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
		out.Log.Println("✅")
	case 1, 2:
		out.Log.Println("🌼 Almost!", sr.Expected())
	default:
		out.Log.Println("🍅 Expected:", sr.Expected())
	}
	alternatives := sr.Alternatives()
	if len(alternatives) > 0 {
		out.Log.Println("Also valid:")
	}
	for _, alt := range alternatives {
		out.Log.Println(answerWithComment(alt))
	}
}

func answerWithComment(crd card.Card) string {
	if text.IsBlank(crd.Comment) {
		return crd.Answer
	}
	return fmt.Sprintf("%s (%s)", crd.Answer, crd.Comment)
}
