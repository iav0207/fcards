package cmd

import (
	. "github.com/iav0207/fcards/internal"
	"github.com/spf13/cobra"
	str "strings"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a flashcard by a given substring",
	Long: `Usage: fcards find 'some phrase I forgot'.
	It will output the found cards contained in the default folder.`,
	Run: runFind,
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func runFind(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		Log.Panicln("Must provide at least one argument - search term.")
	}
	term := args[0]
	found := find(term)
	if countValues(found) == 0 {
		Log.Printf("Term '%s' is not found among %s", term, DefaultTsvFilesPattern)
		return
	}
	printOut(found)
}

func find(term string) map[string][]Card {
	found := make(map[string][]Card)
	for _, path := range AllTsvPaths() {
		for line := range LinesOf(path) {
			if str.Contains(line, term) {
				parsed, err := ParseCard(line)
				Check(err)
				found[path] = append(found[path], *parsed)
			}
		}
	}
	return found
}

func countValues(m map[string][]Card) int {
	count := 0
	for _, cards := range m {
		count += len(cards)
	}
	return count
}

func printOut(occurrences map[string][]Card) {
	for path, cards := range occurrences {
		Log.Printf("%s:", path)
		for i, card := range cards {
			Log.Printf("%d\tQ\t%s", i, card.Question)
			Log.Printf("\tA\t%s", card.Answer)
		}
	}
}
