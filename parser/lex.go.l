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
"(" 
    lval.str = string(y.buf)
    return LBRACK
")" 
    lval.str = string(y.buf)
    return RBRACK
"*"
    lval.str = string(y.buf)
    return MULTI

"AND" 
    return AND
"OR" 
    return OR
"NOT" 
    return NOT

"-h" 
    return HELP
"ui" 
    return UI
"explain" 
    return EXPLAIN
"log" 
    return LOG
"undo" 
    return UNDO
"todo"
    return TODO
"tag" 
    return TAG

"add"|"create" 
    return ADD
"del"|"delete"|"remove" 
    return DELETE
"done"|"complete" 
    return DONE

"age:" 
    return AGE
"due:" 
    return DUE
"loop:" 
    return LOOP
"color:"
    return COLOR
"sort:"
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
