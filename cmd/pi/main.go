package main

import (
	"os"

	"github.com/a-know/pi"
)

func main() {
	os.Exit((&pi.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
