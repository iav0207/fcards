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
	items := []string{"yes", "no"}
	result := UserSelection(prompt, items)
	return str.ToLower(result) == "yes"
}

func UserSelection(prompt string, items []string) string {
	selector := pui.Select{
		Label: prompt,
		Items: items,
	}
	_, result, err := selector.Run()
	FatalIf(err, "Prompt failed")
	return result
}
