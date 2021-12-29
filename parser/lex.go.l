%{
package parser

import (
    "fmt"
    "bufio"
    "os"
    "errors"
    "github.com/SpicyChickenFLY/never-todo-cmd/ast"
)

var (
    src      = bufio.NewReader(os.Stdin)
    buf      []byte
    current  byte
)

type yylexer struct{
    src     *bufio.Reader
    buf     []byte
    empty   bool
    current byte
}

func newLexer(src *bufio.Reader) (y *yylexer) {
    y = &yylexer{src: src}
    y.getc()
    return
}

func (y *yylexer) getc() byte {
    if y.current != 0 {
        y.buf = append(y.buf, y.current)
    }
    y.current = 0
    if b, err := y.src.ReadByte(); err == nil {
        y.current = b
    }
    if debug {
        fmt.Println("getc()->", string(y.current), y.current)
    }
    return y.current
}

func (y yylexer) Error(e string) {
    // fmt.Println(e) 
    ast.ErrorList = append(ast.ErrorList, errors.New("Command not supported: " + e))
    ast.Result = ast.NewRootNode(ast.CMDNotSupport, nil)
}

func (y *yylexer) Lex(lval *yySymType) int {
    // var err error
    c := y.current
    if y.empty {
        c, y.empty = y.getc(), false
    }
%}

%yyc c
%yyn c = y.getc()

setence     \"[^"]*\"|'[^']*'|`[^`]*`
letter      [_a-zA-Z\\]
identifier  {letter}({letter}|{digit})*
white       [ \t\n\r]+
date        {digit}\/{digit}\/{digit}|{digit}\/{digit}
time        (20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d
digit       [0-9]+
importance  ![0-9]

%%

    y.buf = y.buf[:0]

"+" 
    lval.str = string(y.buf)
    return PLUS
"-" 
    lval.str = string(y.buf)
    return MINUS
"*"
    lval.str = string(y.buf)
    return MULTI
"/"
    lval.str = string(y.buf)
    return DIVIDE
"(" 
    lval.str = string(y.buf)
    return LBRACK
")" 
    lval.str = string(y.buf)
    return RBRACK
","
    lval.str = string(y.buf)
    return COMMA

"AND" 
    lval.str = string(y.buf)
    return AND
"OR"
    lval.str = string(y.buf)
    return OR
"NOT" 
    lval.str = string(y.buf)
    return NOT

"-h" 
    return HELP
"ui" 
    lval.str = string(y.buf)
    return UI
"explain" 
    lval.str = string(y.buf)
    return EXPLAIN
"log" 
    lval.str = string(y.buf)
    return LOG
"undo" 
    lval.str = string(y.buf)
    return UNDO

"list"
    lval.str = string(y.buf)
    return LIST
"add"|"create" 
    lval.str = string(y.buf)
    return ADD
"del"|"delete"|"remove" 
    lval.str = string(y.buf)
    return DELETE
"set"
    lval.str = string(y.buf)
    return UPDATE
"todo"
    lval.str = string(y.buf)
    return TODO
"done" 
    lval.str = string(y.buf)
    return DONE
"all"
    lval.str = string(y.buf)
    return ALL

"tag" 
    lval.str = string(y.buf)
    return TAG

"age:" 
    lval.str = string(y.buf)
    return AGE
"due:" 
    lval.str = string(y.buf)
    return DUE
"loop:" 
    lval.str = string(y.buf)
    return LOOP
"color:"
    lval.str = string(y.buf)
    return COLOR
"sort:"
    lval.str = string(y.buf)
    return SORT

{date}
    lval.str = string(y.buf)
    return DATE

{time}
    lval.str = string(y.buf)
    return TIME

{digit}
    lval.str = string(y.buf)
    return NUM

{setence}
    lval.str = string(y.buf[1:len(y.buf)-1])
    return SETENCE

{identifier}
    lval.str = string(y.buf)
    return IDENT

{white} {}

{importance} 
    lval.str = string(y.buf[1:])
    return IMPORTANCE

%%
    y.empty = true
    return int(c)
}
