package cmd

import (
	"fmt"
	"time"
	"math/rand"
	"errors"
	"os"
	fp "path/filepath"
	"github.com/spf13/cobra"
	. "github.com/iav0207/fcards/internal"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a quiz",
	Long: `Start a quiz. Usage: fcards play file1.tsv`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(playCmd)
	directionHelpMsg := fmt.Sprintf("Cards direction. One of: %v", directionValues)
	playCmd.Flags().Var(&directionFlag, "direc", directionHelpMsg)
}

type Direction string
const (
	Straight	Direction = "straight"
	Inverse	 Direction = "inverse"
	Random	  Direction = "random"
)
var directionValues []Direction = []Direction{Straight, Inverse, Random}

func (flag *Direction) String() string {
	return string(*flag)
}

func (flag *Direction) Set(value string) error {
	switch Direction(value) {
	case Straight, Inverse, Random:
		*flag = Direction(value)
		return nil
	default:
		return errors.New(fmt.Sprintf("direction flag value must be one of: %v", directionValues))
	}
}

func (flag *Direction) Type() string {
	return "Direction"
}

var directionFlag Direction

func run(cmd *cobra.Command, args []string) {
	cards := readCards(args)
	fmt.Printf("Read %d cards in total.\n", len(cards))
	fmt.Println("Let's play!")

	shuffle(cards)
	sample := cards[:20]
	applyDirectionFlag(sample)

	reiterate := playRound(sample)
	playRound(reiterate)
}

func readCards(args []string) []Card {
	cards := make([]Card, 0)
	paths := argsOrAllTsvPaths(args)
	for _, filePath := range paths {
		fmt.Printf("Reading cards from file %s\n", filePath)
		cards = append(cards, ReadCards(filePath)...)
	}
	return cards
}

func argsOrAllTsvPaths(args []string) []string {
	if len(args) > 0 {
		return args
	}
	return allTsvPaths()
}

func allTsvPaths() []string {
	pattern := os.Getenv("HOME") + "/.fcards/tsv/*.tsv"
	paths, err := fp.Glob(pattern)
	if err != nil {
		panic(err)
	}
	return paths
}

// Plays a round with given cards and returns those which were given wrong answers to.
func playRound(cards []Card) []Card {
	wrongAnswered := make([]Card, 0)

	for _, card := range cards {
		fmt.Printf("%s:\t", card.Question)
		diff := checkAnswer(ReadLine(), card.Answer)
		if diff > 0 {
			wrongAnswered = append(wrongAnswered, card)
		}
	}

	return wrongAnswered
}

func randomBool() bool {
	return rand.Intn(2) == 0
}

func shuffle(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func (i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func applyDirectionFlag(cards []Card) {
	for i := 0; i < len(cards); i++ {
		if shouldInvert(directionFlag) {
			cards[i].Invert()
		}
	}
}

func shouldInvert(direc Direction) bool {
	return direc == Inverse || (direc == Random && randomBool())
}

func checkAnswer(actual, expected string) int {
	difference := LevenshteinDistance(actual, expected)
	switch difference {
	case 0:
		fmt.Println("✅")
	case 1, 2:
		fmt.Printf("🌼 Almost! %s\n", expected)
	default:
		fmt.Printf("🍅 Expected: %s\n", expected)
	}
	return difference
}

