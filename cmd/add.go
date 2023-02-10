package cmd

import (
	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/data"
	"github.com/iav0207/fcards/internal/in"
	"github.com/iav0207/fcards/internal/model/card"
	"github.com/spf13/cobra"
	"strings"
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
	q := validated(posArgOrUserResponse(args, 0, "Please enter the question (card front side):"))
	a := validated(posArgOrUserResponse(args, 1, "Please enter the answer (card flip side):"))
	path := posArgOrSelection(args, 2, "Where to put it? (file path)", data.AllTsvPaths())

	card := card.New(q, a, "")
	data.AppendToFile(path, card.String())
}

func validated(s string) string {
	check.Require(!strings.ContainsAny(s, "\t\r\n"), "Tabs and line breaks are not allowed")
	return s
}

func posArgOrUserResponse(args []string, pos int, prompt string) string {
	if len(args) > pos {
		return args[pos]
	}
	return in.UserResponse(prompt)
}

func posArgOrSelection(args []string, pos int, prompt string, items []string) string {
	if len(args) > pos {
		return args[pos]
	}
	return in.UserSelection(prompt, items)
}
