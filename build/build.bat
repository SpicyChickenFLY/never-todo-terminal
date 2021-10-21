@ECHO OFF

echo %cd%
echo %~dp0

cd %~dp0

@REM go get modernc.org/goyacc
@REM go get modernc.org/golex

golex -o ..\parser\lex.go ..\parser\lex.go.l
goyacc -o ..\parser\parser.go ..\parser\parser.go.y 

md "%APPDATA%\.nevertodo\"
copy ..\static\data.json %APPDATA%\.nevertodo

@REM go test -v ../parser
@REM go test -v ../controller
@REM go test -v ../data

cd ..

go build -o build\nt.exe
copy build\nt.exe %USERPROFILE%\AppData\Local\Microsoft\WindowsApps
