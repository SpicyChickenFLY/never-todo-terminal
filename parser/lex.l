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

variable @{digit}
setence     \"[^"]*\"|'[^']*'|`[^`]*`
letter      [_a-zA-Z]
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
    lval.str = string(y.buf)
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

"todo"
    lval.str = string(y.buf)
    return TODO
"tag" 
    lval.str = string(y.buf)
    return TAG

"add"|"create" 
    lval.str = string(y.buf)
    return ADD
"del"|"delete"|"remove" 
    lval.str = string(y.buf)
    return DELETE
"done"|"complete" 
    lval.str = string(y.buf)
    return DONE

"age:" 
    return AGE
"due:" 
    return DUE
"loop:" 
    return LOOP

"color:"
    return COLOR

{date}
    lval.str = string(y.buf)
    return DATE

{time}
    lval.str = string(y.buf)
    return TIME

{digit}
    lval.str = string(y.buf)
    return NUM

{variable}
    lval.str = string(y.buf)
    return SETENCE

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