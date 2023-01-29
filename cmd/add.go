package cmd

import (
	. "github.com/iav0207/fcards/internal"
	"github.com/spf13/cobra"
	str "strings"
)

var addCmd = &cobra.Command{
	Use:   `add [question] [answer] [file_path]`,
	Short: "Add a flashcard to the set",
	Long: `All arguments are optional. You will be prompted for any missing value.
	When all positional arguments are specified, the command runs
	non-interactively.`,
	Run: runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// term answer path
func runAdd(cmd *cobra.Command, args []string) {
	q := posArgOrUserResponse(args, 0, "Please enter the question (card front side):")
	validated(q)

	a := posArgOrUserResponse(args, 1, "Please enter the answer (card flip side):")
	validated(a)

	path := posArgOrSelection(args, 2, "Where to put it? (file path)", AllTsvPaths())

	card := NewCard(q, a)
	AppendToFile(path, card.String())
}

func posArgOrUserResponse(args []string, pos int, prompt string) string {
	if len(args) > pos {
		return args[pos]
	}
	return UserResponse(prompt)
}

func posArgOrSelection(args []string, pos int, prompt string, items []string) string {
	if len(args) > pos {
		return args[pos]
	}
	return UserSelection(prompt, items)
}

func validated(s string) string {
	Assert(!str.ContainsAny(s, "\t\r\n"), "Tabs and line breaks are not allowed")
	return s
}
