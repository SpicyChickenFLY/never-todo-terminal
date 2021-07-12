@ECHO OFF

go get modernc.org/goyacc
go get modernc.org/golex

goyacc -o ..\parser\parser.go ..\parser\parser.y 
golex -o ..\parser\lex.go ..\parser\lex.l



go test -v ../parser