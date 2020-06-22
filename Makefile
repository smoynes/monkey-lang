GOPATH := $(PWD)
export GOPATH

DEFAULT: build

run:
	go run monkey

build:
	go build monkey monkey/ast monkey/token monkey/lexer

test:
	go test monkey/lexer monkey/parser

.IGNORE: build test all
