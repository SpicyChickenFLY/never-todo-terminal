package parser

import (
	"bufio"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/ast"
)

var debug = false

// var debug = true

// Parse expose yypase
func Parse(args string) *ast.RootNode {
	yyParse(newLexer(bufio.NewReader(strings.NewReader(args))))
	return ast.Result
}
