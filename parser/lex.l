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
    if debug {
        fmt.Println("getc()->", string(y.current), y.current)
    }
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
setence [\'{3}*\'{3}']
%%

    y.buf = y.buf[:0]

"'" 
    lval.str = string(y.buf)
    return QUOTE
"\"" 
    lval.str = string(y.buf)
    return DQUOTE

":" 
    lval.str = string(y.buf)
    return COLON
"+" 
    lval.str = string(y.buf)
    return PLUS
"-" 
    lval.str = string(y.buf)
    return MINUS

"&" 
    lval.str = string(y.buf)
    return AND
"|" 
    lval.str = string(y.buf)
    return OR
"!" 
    lval.str = string(y.buf)
    return NOT
"^" 
    lval.str = string(y.buf)
    return XOR

"-h" 
    lval.str = string(y.buf)
    return HELP
"ui" 
    lval.str = string(y.buf)
    return UI
"gui" 
    lval.str = string(y.buf)
    return GUI
"explain" 
    lval.str = string(y.buf)
    return EXPLAIN
"log" 
    lval.str = string(y.buf)
    return LOG
"undo" 
    lval.str = string(y.buf)
    return UNDO

"todo"|"task" 
    lval.str = string(y.buf)
    return TASK
"tag" 
    lval.str = string(y.buf)
    return TAG

"add"|"create" 
    lval.str = string(y.buf)
    return ADD
"del"|"delete"|"remove" 
    lval.str = string(y.buf)
    return DELETE
"set"|"edit" 
    lval.str = string(y.buf)
    return SET
"done"|"complete" 
    lval.str = string(y.buf)
    return DONE

"age" 
    lval.str = string(y.buf)
    return AGE

"due" 
    lval.str = string(y.buf)
    return DUE

"like" 
    lval.str = string(y.buf)
    return LIKE

"loop" 
    lval.str = string(y.buf)
    return LOOP

{digit}
    lval.str = string(y.buf)
    return NUM

{identifier}
    lval.str = string(y.buf)
    return IDENT

{white} {}

%%
    y.empty = true
    return int(c)
}