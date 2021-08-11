package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/model"
	"github.com/SpicyChickenFLY/never-todo-cmd/parser"
)

func main() {

	if err := model.Init(""); err != nil {
		fmt.Println("[ERROR]: ", err.Error())
	}
	if err := model.StartTransaction(); err != nil {
		fmt.Println("[ERROR]: ", err.Error())
	}
	// Restore the args to origin command string
	args := strings.Join(os.Args[1:], " ")
	// Parse command string to an AST
	result := parser.Parse(args)
	fmt.Println("[INFO]: parse command string successfully")
	// Execute the AST
	if err := result.Execute(); err != nil {
		fmt.Println("[ERROR]: ", err.Error())
	} else {
		fmt.Println("[INFO]: execute command successfully")
	}
	if err := model.EndTransaction(); err != nil {
		fmt.Println("[ERROR]: ", err.Error())
	}
}
