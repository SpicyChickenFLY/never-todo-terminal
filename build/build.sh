#! /bin/bash

echo $(cd `dirname $0`; pwd)

cd $(cd `dirname $0`; pwd)

golex -o ../parser/lex.go ../parser/lex.l
goyacc -o ../parser/parser.go ../parser/parser.y

mkdir ~/.nevertodo
cp ../static/data.json ~/.nevertodo

cd ..

go build -o build/nt
cp build/nt /usr/bin
