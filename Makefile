GOPATH := $(PWD)
export GOPATH

all: test

build:
	go build monkey/lexer

test:
	go test monkey/lexer

.IGNORE: build test all
