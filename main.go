package main

import (
	"os"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/parser"
	"github.com/SpicyChickenFLY/never-todo-cmd/utils"
)

func main() {

	// Restore the args to origin command string
	// varMap := make(map[string]string, 0)
	// for i := 1; i < len(os.Args); i++ {
	// 	if utils.ContainChinese(os.Args[i]) || len(strings.Split(os.Args[i], " ")) > 1 {
	// 		os.Args[i] = fmt.Sprintf("`%s`", os.Args[i])
	// 	}
	// }
	cmd := strings.Join(os.Args[1:], " ")
	// cmd := "done -11"
	cmdCode := ""
	if len(cmd) > 0 {
		cmdCode = utils.EncodeCmd(cmd)
	}

	// Parse command string to an AST
	result := parser.Parse(cmdCode)

	// Execute the AST
	result.Execute(cmd)
}
