package parser

import (
	"bufio"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/logic"
)

const ( // Command Type
	cmdHelp = 0 + iota
	cmdUI
	cmdGUI
	cmdExplain
)

var debug = false

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
	cmdtype int
}

func (r *root) Execute() {
	switch r.Type() {
	case 0:

	}
}

func (r *root) Type() int {
	return r.cmdtype
}

// StmtNode contain a complex statement
type StmtNode interface {
	Node
	Explain()
}

type help struct {
	root
}

func (h *help) Explain() {
	logic.ShowHelpInfo()
}

// Parse expose yypase
func Parse(args string) {
	yyParse(newLexer(bufio.NewReader(strings.NewReader(args))))
}
