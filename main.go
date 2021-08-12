package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/parser"
)

func main() {
	if err := model.Init("./static/data.json"); err != nil {
		panic(err)
	}
	if err := model.Begin(); err != nil {
		panic(err)
	}
	// Restore the args to origin command string
	for i, arg := range os.Args[1:] {
		if len(strings.Split(arg, " ")) > 1 {
			os.Args[i] = fmt.Sprintf("`%s`", arg)
		}
	}
	cmd := strings.Join(os.Args[1:], " ")
	// fmt.Println("[INFO]: ", cmd)
	// Parse command string to an AST
	result := parser.Parse(cmd)
	// fmt.Println("[INFO]: parse command string successfully")
	// Execute the AST
	if err := result.Execute(); err != nil {
		model.RollBack()
	} else {
		if err := model.Commit(); err != nil {
			panic(err)
		}
	}
}
