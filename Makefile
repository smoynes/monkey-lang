GOPATH := $(PWD)
export GOPATH

DEFAULT: build

run:
	go run monkey

build:
	go build monkey monkey/ast monkey/evaluator monkey/lexer monkey/object monkey/parser monkey/repl monkey/token

test:
	go test monkey/lexer monkey/parser monkey/ast monkey/object monkey/evaluator

.IGNORE: build test all
