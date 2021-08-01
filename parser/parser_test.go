package parser

import (
	"fmt"
	"testing"

	"github.com/SpicyChickenFLY/never-todo-cmd/parser/ast"
)

type testCase struct {
	input      string
	output     ast.Node
	supposeErr error
}

func TestParser(t *testing.T) {
	testcases := []string{
		// help
		"-h",
		"todo add -h",
		"todo add me -h",
		"tag -h",
		// task
		"todo",
		"look for me",
		"todo add remeber to go shopping",
		"add 'del: 1 + 1 & 2 '",
		"explain todo done 1-4 20-15",
		"del 20 2-5",
		//
		"tag 9",
		// err
	}
	for _, tc := range testcases {
		fmt.Printf("testcase:%s\n", tc)
		Parse(tc)
		fmt.Println(result)
		result.Explain()
		// if err := result.Execute(); err != nil {
		// 	fmt.Println(err.Error())
		// }
		fmt.Println("")
	}
}
