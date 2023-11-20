package main

import (
	"os"

	"github.com/gofofoca/gofofoca/internal/gofofoca/command"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
