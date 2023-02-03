package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fcards",
	Short: "Simple command-line flashcards",
	Long:  `Point it at a tab-separated file and play a quiz.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defaultArgsIfNoneGiven()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var defaultArgs []string = []string{"play", "--direc", "random"}

func defaultArgsIfNoneGiven() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, defaultArgs...)
	}
}

func init() {
}
