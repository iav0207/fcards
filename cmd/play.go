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

type Direction string
const (
    Straight    Direction = "straight"
    Inverse     Direction = "inverse"
    Random      Direction = "random"
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
    cards := make([]Card, 0)
    paths := argsOrAllTsvPaths(args)
    for _, filePath := range paths {
        fmt.Printf("Reading cards from file %s\n", filePath)
        cards = append(cards, ReadCards(filePath)...)
    }
    fmt.Printf("Read %d cards in total.\n", len(cards))
    fmt.Println("Let's play!")

    shuffle(cards)

    for _, card := range cards[:20] {
        if shouldInvert(directionFlag) {
            card.Invert()
        }
        fmt.Printf("%s:\t", card.Question)
        checkAnswer(ReadLine(), card.Answer)
    }
}

func argsOrAllTsvPaths(args []string) []string {
    if len(args) > 0 {
        return args
    }
    return allTsvPaths()
}

func allTsvPaths() []string {
    pattern := fmt.Sprintf("%s/.fcards/tsv/*.tsv", os.Getenv("HOME"))
    paths, err := fp.Glob(pattern)
    if err != nil {
        panic(err)
    }
    return paths
}

func shouldInvert(direc Direction) bool {
    return direc == Inverse || (direc == Random && randomBool())
}

func randomBool() bool {
    return rand.Intn(2) == 0
}

func shuffle(cards []Card) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(cards), func (i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func checkAnswer(actual, expected string) {
    difference := LevenshteinDistance(actual, expected)
    switch difference {
    case 0:
        fmt.Println("✅")
    case 1, 2:
        fmt.Printf("🌼 Almost! %s\n", expected)
    default:
        fmt.Printf("🍅 Expected: %s\n", expected)
    }
}

func init() {
	rootCmd.AddCommand(playCmd)
    directionHelpMsg := fmt.Sprintf("Cards direction. One of: %v", directionValues)
    playCmd.Flags().Var(&directionFlag, "direc", directionHelpMsg)
}
