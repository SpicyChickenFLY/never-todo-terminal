#! /bin/bash

echo $(cd `dirname $0`; pwd)

cd $(cd `dirname $0`; pwd)

golex -o ../parser/lex.go ../parser/lex.go.l
goyacc -o ../parser/parser.go ../parser/parser.go.y

mkdir ~/.nevertodo
cp ../static/data_tmpl.json ~/.nevertodo/data.json

cd ..

go build -o build/nt
cp build/nt /usr/bin
