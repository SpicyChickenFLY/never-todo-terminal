package parser

import (
	"bufio"
	"strings"

	"github.com/SpicyChickenFLY/never-todo-cmd/ast"
)

var debug = false

// var debug = true

// Parse expose yypase
func Parse(stmt string) *ast.RootNode {
	yyParse(newLexer(bufio.NewReader(strings.NewReader(stmt))))
	return ast.Result
}

// // Parse expose yypase
// func Parse(stmt string, varMap map[string]string) *ast.RootNode {
// 	ast.SetVarMap(varMap)
// 	yyParse(newLexer(bufio.NewReader(strings.NewReader(stmt))))
// 	return ast.Result
// }
