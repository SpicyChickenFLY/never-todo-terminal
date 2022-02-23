package parser

const (
	Num = iota
	Ident
	Assign
	Unassign
	Importance
	Age
	Due
	Loop
	IDGroup
)

func lex(str string) (int, interface{}) {
	return Ident, str
}
