package main

import (
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/parser"
	"github.com/SpicyChickenFLY/never-todo-cmd/pkgs/utils"
)

func main() {

	// Restore the args to origin command string
	// varMap := make(map[string]string, 0)
	// for i := 1; i < len(os.Args); i++ {
	// 	if utils.ContainChinese(os.Args[i]) || len(strings.Split(os.Args[i], " ")) > 1 {
	// 		os.Args[i] = fmt.Sprintf("`%s`", os.Args[i])
	// 	}
	// }
	// fmt.Println(os.Args)
	cmd := strings.Join(os.Args[1:], " ")

	// cmd := "list"
	if len(cmd) > 0 {
		cmd = utils.EncodeCmd(cmd)
	}

	// Parse command string to an AST
	// result := parser.Parse(cmd)

	result := parser.ParseRoot(os.Args[1:])

	// Execute the AST
	result.Execute(cmd)
}
