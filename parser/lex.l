%{
package parser

import (
    "fmt"
    "bufio"
    "log"
    "os"
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
    fmt.Println("getc()->", string(y.current))
    return y.current
}

func (y yylexer) Error(e string) {
    log.Fatal(e)
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

digit  [0-9]+
letter        [_a-zA-Z]
identifier    {letter}({letter}|{digit})*
white		[ \t\n\r]+

%%

":" return COLON
"+" return PLUS
"-" return MINUS

"&" return AND
"|" return OR
"!" return NOT
"^" return XOR

"help" return HELP
"ui" return UI
"gui" return GUI
"explain" return EXPLAIN
"log" return LOG
"undo" return UNDO

"todo"|"task" return TASK
"tag" return TAG

"add"|"create" return ADD
"del"|"delete"|"remove" return DELETE
"set"|"edit" return SET
"done"|"complete" return DONE
"age" return AGE
"due" return DUE
"like" return LIKE
"loop" return LOOP

{digit} return NUM
{identifier} return IDENT
{white} return 0

%%
    y.empty = true
    return int(c)
}