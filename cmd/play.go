package cmd

import (
	"fmt"
    "time"
    "math/rand"
	"github.com/spf13/cobra"
    "github.com/iav0207/fcards/internal"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a quiz",
    Long: `Start a quiz. Usage: fcards play file1.tsv`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
    // TODO if args is empty, traverse all tsv files in ~/.fcards
    cards := make([]internal.Card, 0)
    for _, filePath := range args {
        fmt.Printf("Reading cards from file %s\n", filePath)
        cards = append(cards, internal.ReadCards(filePath)...)
    }
    fmt.Printf("Read %d cards in total.\n", len(cards))
    fmt.Println("Let's play!")

    shuffle(cards)

    for _, card := range cards[:20] {
        fmt.Printf("%s:\t", card.Question)
        checkAnswer(internal.ReadLine(), card.Answer)
        // TODO offer to update the card?
    }
}

func shuffle(cards []internal.Card) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(cards), func (i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func checkAnswer(actual, expected string) {
    correct := internal.Compare(actual, expected) == 0
    if correct {
        fmt.Println("✅")
    } else {
        fmt.Printf("🍅 Expected: %s\n", expected)
    }
}

func init() {
	rootCmd.AddCommand(playCmd)
    // TODO a flag to randomly invert some cards
}
