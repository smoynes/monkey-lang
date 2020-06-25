package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
)

var interactive = flag.Bool("i", false, "interactive mode")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-i] [<filename>]\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	if *interactive {
		startRepl()
	} else {
		if flag.NArg() != 1 {
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
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [<filename>]\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}

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
