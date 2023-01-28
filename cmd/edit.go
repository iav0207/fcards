package cmd

import (
	. "github.com/iav0207/fcards/internal"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
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
	Assert(len(found) == 1, "More than one occurrence found. Please make the request more specific")

	path, card := firstCard(found)

	if UserConfirms("Do you want to edit the found card?") {
		edit(path, card)
	}
}

func edit(path string, card Card) {
	q := UserResponse("Please enter the new question (card front side) below:")
	a := UserResponse("Please enter the new answer (card flip side) below:")
	content := make([]string, 0)
	for line := range LinesOf(path) {
		if line == card.String() {
			line = NewCard(q, a).String()
		}
		content = append(content, line)
	}
	OverwriteFileWithLines(path, content)
}

func firstCard(m map[string][]Card) (string, Card) {
	for k, v := range m {
		return k, v[0]
	}
	panic("")
}