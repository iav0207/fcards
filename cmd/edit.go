package cmd

import (
	"fmt"
	. "github.com/iav0207/fcards/internal"
	"github.com/iav0207/fcards/internal/model"
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
	Assert(len(found) > 0)
	Require(countValues(found) == 1, "More than one occurrence found. Please make the request more specific")

	path, card := firstCard(found)

	if UserConfirms("Do you want to edit the found card?") {
		edit(path, card)
	}
}

func edit(path string, card model.Card) {
	q := defaultedInput("the new question (card front side)", card.Question)
	a := defaultedInput("the new answer (card flip side)", card.Answer)
	content := make([]string, 0)
	updatedLines := 0
	for line := range LinesFrom(path) {
		if line == card.String() { // FIXME support comments
			line = model.NewCard(q, a, "").String()
			updatedLines++
		}
		content = append(content, line)
	}
	Assert(updatedLines == 1, "Expected to update one line, was about to update", updatedLines)
	OverwriteFileWithLines(path, content)
}

func defaultedInput(ofWhat, defaultValue string) string {
	promptFmt := "Please enter %s below. Leave blank to keep it as is."
	prompt := fmt.Sprintf(promptFmt, ofWhat)
	input := UserResponse(prompt)
	if IsBlank(input) {
		return defaultValue
	}
	return input
}

func firstCard(m map[string][]model.Card) (string, model.Card) {
	for k, v := range m {
		return k, v[0]
	}
	panic("")
}
