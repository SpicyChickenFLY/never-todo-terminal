package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/parser"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

func main() {
	if err := model.Init("./static/data.json"); err != nil {
		panic(err)
	}
	if err := model.Begin(); err != nil {
		panic(err)
	}

	// Restore the args to origin command string
	// varMap := make(map[string]string, 0)
	for i := 1; i < len(os.Args); i++ {
		if utils.ContainChinese(os.Args[i]) || len(strings.Split(os.Args[i], " ")) > 1 {
			os.Args[i] = fmt.Sprintf("`%s`", os.Args[i])
		}
	}
	cmd := strings.Join(os.Args[1:], " ")

	// Parse command string to an AST
	result := parser.Parse(cmd)

	// Execute the AST
	if err := result.Execute(cmd); err != nil {
		model.RollBack()
	} else {
		if err := model.Commit(); err != nil {
			panic(err)
		}
	}
}
