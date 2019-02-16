package main

import (
	"fmt"
	"os"

	"github.com/a-know/pi"
)

func main() {
	fmt.Println("Hello, pi!!")
	os.Exit((&pi.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
