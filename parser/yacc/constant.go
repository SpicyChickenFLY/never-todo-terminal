package parser

import (
	"bufio"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/parser/ast"
)

var debug = false

var result ast.Node

// Parse expose yypase
func Parse(args string) ast.Node {
	yyParse(newLexer(bufio.NewReader(strings.NewReader(args))))
	return result
}
