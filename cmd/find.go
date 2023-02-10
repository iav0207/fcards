package cmd

import (
	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/data"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/out"
	"github.com/iav0207/fcards/internal/text"
	"github.com/spf13/cobra"
	"strings"
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

func runFindReturnFound(args []string) map[string][]card.Card {
	check.Require(len(args) > 0, "Must provide at least one argument - search term.")
	term := strings.Join(args, " ")
	found := find(term)
	check.Require(countValues(found) > 0, "Term '%s' is not found among %s", term, data.DefaultTsvFilesPattern)
	printOut(found)
	return found
}

func find(term string) map[string][]card.Card {
	found := make(map[string][]card.Card)
	for _, path := range data.AllTsvPaths() {
		for line := range data.LinesFrom(path) {
			if strings.Contains(line, term) {
				parsed, err := data.ParseCard(line)
				check.PanicIf(err)
				found[path] = append(found[path], *parsed)
			}
		}
	}
	return found
}

func countValues(m map[string][]card.Card) int {
	count := 0
	for _, cards := range m {
		count += len(cards)
	}
	return count
}

func printOut(occurrences map[string][]card.Card) {
	for path, cards := range occurrences {
		for _, c := range cards {
			out.Log.Println(text.TabSeparated(path, c.Question, c.Answer))
		}
	}
}
