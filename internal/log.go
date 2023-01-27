package internal

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "", 0)
