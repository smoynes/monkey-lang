package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
)

var interactive = flag.Bool("i", false, "interactive REPL")

func main() {
	flag.Parse()

	if *interactive {
		startRepl()
	} else {
		if flag.NArg() == 0 {
			os.Stderr.WriteString("no filename")
			flag.Usage()
		}
		startCommand(flag.Arg(0))
	}
}

func startRepl() {
	fmt.Printf("Monkey language interpreter.\n")
	repl.Start(os.Stdin, os.Stdout)
}

func startCommand(filename string) {
	inputFile, err := os.Open(filename)
	if err != nil {
		os.Stderr.WriteString(filename)
		return
	}
	defer inputFile.Close()

	input, err := ioutil.ReadAll(inputFile)
	env := object.NewEnvironment()
	lexer := lexer.New(string(input))
	parser := parser.New(lexer)

	program := parser.ParseProgram()
	if len(parser.Errors()) > 0 {
		printParseErrors(os.Stdout, parser.Errors())
		return
	}

	evaluator.Eval(program, env)
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
