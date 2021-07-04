package parser

import "log"

type Tokenizer struct {
	query   Query
	scanner *Scanner
}

func (tkn *Tokenizer) Lex(lval *yySymType) int {
	var typ int
	var val string

	for {
		typ, _, val = tkn.scanner.Scan()
		if typ == EOF {
			return 0
		}
		if typ != WS {
			break
		}
	}
	lval.str = val
	return typ
}
func (tkn *Tokenizer) Error(err string) {
	log.Fatal(err)
}
