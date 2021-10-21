package parser

import (
	"fmt"
	"testing"

	"github.com/SpicyChickenFLY/never-todo-cmd/ast"
)

type testCase struct {
	input      string
	output     ast.RootNode
	supposeErr error
}

func TestParser(t *testing.T) {
	testcases := []string{
		// help
		"-h",
		"add -h",
		"todo add -h",
		"todo add me -h",
		"tag -h",
		// task
		"todo 29",
		"NOT('A' AND \"B\" OR `C`)",
		"todo age:2020/03/02-2021/02/03",
		"add play ball +exercise !3",
		"add 'del:  1 + 1 & 2 ' due:2020/04/3",
		"add `你好`",
		"todo add go shopping due:2015/03/12",
		"done 1-4 20-15",
		"del 20 2-5",
		//
		"tag 9",
		"tag add oweh",
		// err
	}
	for _, tc := range testcases {
		fmt.Printf("testcase: %s\n", tc)
		Parse(tc)
		ast.Result.Explain()
		fmt.Println("")
	}
}
