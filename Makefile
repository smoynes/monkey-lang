GOPATH := $(PWD)
export GOPATH

all: test

build:
	go build lexer token

test:
	go test lexer

.IGNORE: build test all
