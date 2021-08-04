@ECHO OFF

echo %cd%
echo %~dp0

cd %~dp0

go get modernc.org/goyacc
go get modernc.org/golex

goyacc -o ..\parser\parser.go ..\parser\parser.y 
golex -o ..\parser\lex.go ..\parser\lex.l

go test -v ../parser
@REM go test -v ../controller
@REM go test -v ../data

cd ..

go build -o never.exe