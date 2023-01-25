package internal

import (
	"bufio"
	"os"
)

func ReadLine() string {
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return line
}
