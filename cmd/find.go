package cmd

import (
	. "github.com/iav0207/fcards/internal"
	"github.com/spf13/cobra"
	str "strings"
)

var findCmd = &cobra.Command{
	Use:   "find [search_term]",
	Short: "Find a flashcard by a given substring",
	Long:  `It will output the found cards contained in the default folder.`,
	Run:   func(cmd *cobra.Command, args []string) { runFindReturnFound(args) },
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func runFindReturnFound(args []string) map[string][]Card {
	Require(len(args) > 0, "Must provide at least one argument - search term.")
	term := str.Join(args, " ")
	found := find(term)
	Require(countValues(found) > 0, "Term '%s' is not found among %s", term, DefaultTsvFilesPattern)
	printOut(found)
	return found
}

func find(term string) map[string][]Card {
	found := make(map[string][]Card)
	for _, path := range AllTsvPaths() {
		for line := range LinesOf(path) {
			if str.Contains(line, term) {
				parsed, err := ParseCard(line)
				PanicIf(err)
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
		for _, card := range cards {
			Log.Println(TabSeparated(path, card.Question, card.Answer))
		}
	}
}
