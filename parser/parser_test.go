package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

type Command struct{}

func TestParser(t *testing.T) {
	testcases := []string{
		"help",
		"todo",
		"todo add hi",
		"tag 9",
	}
	for _, tc := range testcases {
		fmt.Println("testcase: ", tc)
		yyParse(newLexer(bufio.NewReader(strings.NewReader(tc))))
		fmt.Println("")
	}
}
