package internal

import (
	"bufio"
	pui "github.com/manifoldco/promptui"
	"os"
	str "strings"
)

func UserResponse(prompt string) string {
	Log.Println(prompt)
	return ReadLine()
}

func ReadLine() string {
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return trimRightLineBreak(line)
}

func UserConfirms(prompt string) bool {
	selector := pui.Select{
		Label: prompt + " [y/n]",
		Items: []string{"y", "n"},
	}
	_, result, err := selector.Run()
	if err != nil {
		Log.Fatalf("Prompt failed %w\n", err)
	}
	return str.ToLower(result) == "y"
}
