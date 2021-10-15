// Code generated by golex. DO NOT EDIT.

package parser

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/SpicyChickenFLY/never-todo-cmd/ast"
	"os"
)

var (
	src     = bufio.NewReader(os.Stdin)
	buf     []byte
	current byte
)

type yylexer struct {
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
	ast.ErrorList = append(ast.ErrorList, errors.New("Command not supported: "+e))
	ast.Result = ast.NewRootNode(ast.CMDNotSupport, nil)
}

func (y *yylexer) Lex(lval *yySymType) int {
	// var err error
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.buf = y.buf[:0]

	goto yystart1

yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '(':
		goto yystate8
	case c == ')':
		goto yystate9
	case c == '*':
		goto yystate10
	case c == '+':
		goto yystate11
	case c == '-':
		goto yystate12
	case c == '0' || c == '1':
		goto yystate14
	case c == '2':
		goto yystate27
	case c == 'A':
		goto yystate29
	case c == 'N':
		goto yystate33
	case c == 'O':
		goto yystate36
	case c == '\'':
		goto yystate7
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c == '`':
		goto yystate38
	case c == 'a':
		goto yystate39
	case c == 'c':
		goto yystate45
	case c == 'd':
		goto yystate61
	case c == 'e':
		goto yystate71
	case c == 'l':
		goto yystate78
	case c == 'r':
		goto yystate84
	case c == 's':
		goto yystate88
	case c == 't':
		goto yystate93
	case c == 'u':
		goto yystate99
	case c >= '3' && c <= '9':
		goto yystate19
	case c >= 'B' && c <= 'M' || c >= 'P' && c <= 'Z' || c == '\\' || c == '_' || c == 'b' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'q' || c >= 'v' && c <= 'z':
		goto yystate30
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule29
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate4
	}

yystate4:
	c = y.getc()
	goto yyrule30

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c >= '\x01' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate5
	}

yystate6:
	c = y.getc()
	goto yyrule27

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate6
	case c >= '\x01' && c <= '&' || c >= '(' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = y.getc()
	goto yyrule3

yystate9:
	c = y.getc()
	goto yyrule4

yystate10:
	c = y.getc()
	goto yyrule5

yystate11:
	c = y.getc()
	goto yyrule1

yystate12:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == 'h':
		goto yystate13
	}

yystate13:
	c = y.getc()
	goto yyrule9

yystate14:
	c = y.getc()
	switch {
	default:
		goto yyrule26
	case c == '/':
		goto yystate15
	case c == 'd':
		goto yystate20
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate15:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate16
	}

yystate16:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c == '/':
		goto yystate17
	case c >= '0' && c <= '9':
		goto yystate16
	}

yystate17:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate18:
	c = y.getc()
	switch {
	default:
		goto yyrule24
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyrule26
	case c == '/':
		goto yystate15
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate20:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate21
	}

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '5':
		goto yystate22
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate23
	}

yystate23:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate24
	}

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '5':
		goto yystate25
	}

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate26
	}

yystate26:
	c = y.getc()
	goto yyrule25

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule26
	case c == '/':
		goto yystate15
	case c >= '0' && c <= '3':
		goto yystate28
	case c >= '4' && c <= '9':
		goto yystate19
	}

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule26
	case c == '/':
		goto yystate15
	case c == ':':
		goto yystate21
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'N':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate30:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'D':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate32:
	c = y.getc()
	switch {
	default:
		goto yyrule6
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'O':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'T':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate35:
	c = y.getc()
	switch {
	default:
		goto yyrule8
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate36:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'R':
		goto yystate37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate37:
	c = y.getc()
	switch {
	default:
		goto yyrule7
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate38:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '`':
		goto yystate6
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate38
	}

yystate39:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'd':
		goto yystate40
	case c == 'g':
		goto yystate42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'c' || c == 'e' || c == 'f' || c >= 'h' && c <= 'z':
		goto yystate30
	}

yystate40:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'd':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate30
	}

yystate41:
	c = y.getc()
	switch {
	default:
		goto yyrule16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate42:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate43:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == ':':
		goto yystate44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate44:
	c = y.getc()
	goto yyrule19

yystate45:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate46
	case c == 'r':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c == 'p' || c == 'q' || c >= 's' && c <= 'z':
		goto yystate30
	}

yystate46:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate47
	case c == 'm':
		goto yystate51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'k' || c >= 'n' && c <= 'z':
		goto yystate30
	}

yystate47:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate48:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'r':
		goto yystate49
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate30
	}

yystate49:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == ':':
		goto yystate50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate50:
	c = y.getc()
	goto yyrule22

yystate51:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'p':
		goto yystate52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate30
	}

yystate52:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate30
	}

yystate53:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate54
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate54:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate30
	}

yystate55:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate56:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate57:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate58:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate59
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate30
	}

yystate59:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate60
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate30
	}

yystate60:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate61:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate62
	case c == 'o':
		goto yystate67
	case c == 'u':
		goto yystate68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'n' || c >= 'p' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate30
	}

yystate62:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate63
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate30
	}

yystate63:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate64:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate30
	}

yystate65:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate66:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate67:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate30
	}

yystate68:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate69:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == ':':
		goto yystate70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate70:
	c = y.getc()
	goto yyrule20

yystate71:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'x':
		goto yystate72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate30
	}

yystate72:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'p':
		goto yystate73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate30
	}

yystate73:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate30
	}

yystate74:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate75
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate30
	}

yystate75:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'i':
		goto yystate76
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate30
	}

yystate76:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate77
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate30
	}

yystate77:
	c = y.getc()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate78:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate79
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate79:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'g':
		goto yystate80
	case c == 'o':
		goto yystate81
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate80:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate81:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'p':
		goto yystate82
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate30
	}

yystate82:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == ':':
		goto yystate83
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate83:
	c = y.getc()
	goto yyrule21

yystate84:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate30
	}

yystate85:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'm':
		goto yystate86
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate30
	}

yystate86:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate87
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate87:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'v':
		goto yystate65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'u' || c >= 'w' && c <= 'z':
		goto yystate30
	}

yystate88:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate89
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate89:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'r':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate30
	}

yystate90:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate91
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate30
	}

yystate91:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == ':':
		goto yystate92
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate92:
	c = y.getc()
	goto yyrule23

yystate93:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate94
	case c == 'o':
		goto yystate96
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'b' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate94:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'g':
		goto yystate95
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate30
	}

yystate95:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate96:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'd':
		goto yystate97
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate30
	}

yystate97:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate98
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate98:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate99:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'i':
		goto yystate100
	case c == 'n':
		goto yystate101
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate30
	}

yystate100:
	c = y.getc()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yystate101:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'd':
		goto yystate102
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate30
	}

yystate102:
	c = y.getc()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate103
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate30
	}

yystate103:
	c = y.getc()
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '\\' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate30
	}

yyrule1: // "+"
	{

		lval.str = string(y.buf)
		return PLUS
	}
yyrule2: // "-"
	{

		lval.str = string(y.buf)
		return MINUS
	}
yyrule3: // "("
	{

		lval.str = string(y.buf)
		return LBRACK
	}
yyrule4: // ")"
	{

		lval.str = string(y.buf)
		return RBRACK
	}
yyrule5: // "*"
	{

		lval.str = string(y.buf)
		return MULTI
	}
yyrule6: // "AND"
	{

		return AND
	}
yyrule7: // "OR"
	{

		return OR
	}
yyrule8: // "NOT"
	{

		return NOT
	}
yyrule9: // "-h"
	{

		return HELP
	}
yyrule10: // "ui"
	{

		return UI
	}
yyrule11: // "explain"
	{

		return EXPLAIN
	}
yyrule12: // "log"
	{

		return LOG
	}
yyrule13: // "undo"
	{

		return UNDO
	}
yyrule14: // "todo"
	{

		return TODO
	}
yyrule15: // "tag"
	{

		return TAG
	}
yyrule16: // "add"|"create"
	{

		return ADD
	}
yyrule17: // "del"|"delete"|"remove"
	{

		return DELETE
	}
yyrule18: // "done"|"complete"
	{

		return DONE
	}
yyrule19: // "age:"
	{

		return AGE
	}
yyrule20: // "due:"
	{

		return DUE
	}
yyrule21: // "loop:"
	{

		return LOOP
	}
yyrule22: // "color:"
	{

		return COLOR
	}
yyrule23: // "sort:"
	{

		return SORT
	}
yyrule24: // {date}
	{

		lval.str = string(y.buf)
		return DATE
	}
yyrule25: // {time}
	{

		lval.str = string(y.buf)
		return TIME
	}
yyrule26: // {digit}
	{

		lval.str = string(y.buf)
		return NUM
	}
yyrule27: // {setence}
	{

		lval.str = string(y.buf[1 : len(y.buf)-1])
		return SETENCE
	}
yyrule28: // {identifier}
	{

		lval.str = string(y.buf)
		return IDENT
	}
yyrule29: // {white}
	{
		{
		}
		goto yystate0
	}
yyrule30: // {importance}
	if true { // avoid go vet determining the below panic will not be reached

		lval.str = string(y.buf[1:])
		return IMPORTANCE
	}
	panic("unreachable")

yyabort: // no lexem recognized
	//
	// silence unused label errors for build and satisfy go vet reachability analysis
	//
	{
		if false {
			goto yyabort
		}
		if false {
			goto yystate0
		}
		if false {
			goto yystate1
		}
	}

	y.empty = true
	return int(c)
}
