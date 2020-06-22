GOPATH := $(PWD)
export GOPATH

DEFAULT: build

run:
	go run monkey

build:
	go build monkey monkey/ast monkey/token monkey/lexer monkey/parser

test:
	go test monkey/lexer monkey/parser monkey/ast

.IGNORE: build test all
