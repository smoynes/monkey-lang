package main

import (
	"fmt"
	"os"
	"monkey/repl"
)

func main() {
	fmt.Printf("Monkey language interpreter.\n")
	repl.Start(os.Stdin, os.Stdout)
}
