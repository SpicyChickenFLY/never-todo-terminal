package main

import (
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/parser"
)

func main() {
	// cmd.Execute()
	args := strings.Join(os.Args[1:], " ")
	parser.Parse(args)
}
