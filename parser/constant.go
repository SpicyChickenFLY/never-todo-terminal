package parser

import (
	"bufio"
	"strings"
)

const ( // Command Type
	help = 0 + itoa,
	ui,
	gui,
	explain,
)

// Node is the basic element of the AST
type Node interface {
	Text() string
	SetText(text string)
}

// RootNode is the root of the AST
type RootNode interface {
	Execute()
	Type()
}

type root struct {
}

func (r *root) Execute() {
	switch r.Type() {
	case "":

	}
}

// StmtNode contain a complex statement
type StmtNode interface {
	Node
	Explain()
}

// Parse expose yypase
func Parse(args string) {
	yyParse(newLexer(bufio.NewReader(strings.NewReader(args))))
}
