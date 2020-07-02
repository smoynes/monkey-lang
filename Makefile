DEFAULT: build

all: test build

run:
	go run monkey -i

build:
	go build ./cmd/monkey

test:
	go test -v ./lexer ./parser ./ast ./object ./evaluator 

.IGNORE: run build test all
