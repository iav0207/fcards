package cmd

import (
	"fmt"
	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/data"
	"github.com/iav0207/fcards/internal/in"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/iav0207/fcards/internal/text"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [search_term]",
	Short: "Edit a flashcard by a given substring",
	Long: `Usage: fcards edit 'phrase I want to correct'.
	It will search the term among the cards contained in the default folder,
	then will ask to refedine the card.`,
	Run: runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func runEdit(cmd *cobra.Command, args []string) {
	found := runFindReturnFound(args)
	check.Assert(len(found) > 0)
	check.Require(countValues(found) == 1, "More than one occurrence found. Please make the request more specific")

	path, card := firstCard(found)

	if in.UserConfirms("Do you want to edit the found card?") {
		edit(path, card)
	}
}

func edit(path string, c card.Card) {
	q := defaultedInput("the new question (card front side)", c.Question)
	a := defaultedInput("the new answer (card flip side)", c.Answer)
	content := make([]string, 0)
	updatedLines := 0
	cRaw := c.String()
	for line := range data.LinesFrom(path) {
		if line == cRaw {
			line = card.New(q, a, c.Comment).String()
			updatedLines++
		}
		content = append(content, line)
	}
	check.Assert(updatedLines == 1, "Expected to update one line, was about to update", updatedLines)
	data.OverwriteFileWithLines(path, content)
}

func defaultedInput(ofWhat, defaultValue string) string {
	promptFmt := "Please enter %s below. Leave blank to keep it as is."
	prompt := fmt.Sprintf(promptFmt, ofWhat)
	input := in.UserResponse(prompt)
	if text.IsBlank(input) {
		return defaultValue
	}
	return input
}

func firstCard(m map[string][]card.Card) (string, card.Card) {
	for k, v := range m {
		return k, v[0]
	}
	panic("")
}
