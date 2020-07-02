package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/smoynes/monkey-lang/lexer"
	"github.com/smoynes/monkey-lang/parser"
	"github.com/smoynes/monkey-lang/evaluator"
	"github.com/smoynes/monkey-lang/object"
)

const prompt = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()

		// Exit if end of input or error.
		if !scanned {
			return
		}

		// Create a parser for each line.
		line := scanner.Text()
		lexer := lexer.New(line)
		parser := parser.New(lexer)

		program := parser.ParseProgram()
		if len(parser.Errors()) > 0 {
			printParseErrors(out, parser.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
