GOPATH := $(PWD)
export GOPATH

all: test

run:
	go run monkey

build:
	go build monkey/lexer monkey/token monkey/repl monkey

test:
	go test monkey/lexer

.IGNORE: build test all
