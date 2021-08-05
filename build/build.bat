@ECHO OFF

echo %cd%
echo %~dp0

cd %~dp0

@REM go get modernc.org/goyacc
@REM go get modernc.org/golex

golex -o ..\parser\lex.go ..\parser\lex.l
goyacc -o ..\parser\parser.go ..\parser\parser.y 


go test -v ../parser
@REM go test -v ../controller
@REM go test -v ../data

cd ..

go build -o never.exe