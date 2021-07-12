package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	input      string
	output     RootNode
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
		"todo done 1-4 20",
		"del 20 2-5",
		//
		"tag 9",
		// err
	}
	for _, tc := range testcases {
		fmt.Printf("testcase:%s\n", tc)
		yyParse(newLexer(bufio.NewReader(strings.NewReader(tc))))
		fmt.Println("")
	}
}
