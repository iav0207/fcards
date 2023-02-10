package in

import (
	"bufio"
	pui "github.com/manifoldco/promptui"
	"os"
	"runtime"
	"strings"

	"github.com/iav0207/fcards/internal/check"
	"github.com/iav0207/fcards/internal/out"
)

func UserResponse(prompt string) string {
	out.Log.Println(prompt)
	return ReadLine()
}

func ReadLine() string {
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(line, lineBreak())
}

func UserConfirms(prompt string) bool {
	items := []string{"yes", "no"}
	result := UserSelection(prompt, items)
	return strings.ToLower(result) == "yes"
}

func UserSelection(prompt string, items []string) string {
	selector := pui.Select{
		Label: prompt,
		Items: items,
	}
	_, result, err := selector.Run()
	check.FatalIf(err, "Prompt failed")
	return result
}

func lineBreak() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}
