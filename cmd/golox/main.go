package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atmask/golox/internal/lox"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		if err := runFile(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(64)
		}
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	errs := run(string(bytes))
	if len(errs) > 0 {
		os.Exit(65)
	}

	return nil
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		// Igore Errors and don't poison the REPL
		run(scanner.Text())

	}
}

func run(source string) []lox.ScanError {
	scanner := lox.NewScanner(source)
	tokens, errs := scanner.ScanTokens()

	for _, e := range errs {
		fmt.Fprintln(os.Stderr, e.Error())
	}

	for _, t := range tokens {
		fmt.Println(t)
	}

	return errs
}
