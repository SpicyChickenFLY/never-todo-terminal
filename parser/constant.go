package parser

import (
	"bufio"
	"strings"
)

const ( // Command Type

)

type Commander interface {
	Execute()
	Explain()
}

type CommandTree struct {
	Type int
}

func NewCommandTree() CommandTree {
	return CommandTree{}
}

type Node struct {
}

func Parse(args string) {
	yyParse(newLexer(bufio.NewReader(strings.NewReader(args))))
}
