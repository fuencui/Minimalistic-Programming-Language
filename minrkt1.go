package main

import (
	"bufio"
	"fmt"
	"os"

	"my.com/cs5400/minrkt"
)

func main() {
	fmt.Println("Welcome to minimalistic racket phase!")
	scanner := bufio.NewScanner(os.Stdin)
	defineNameMap := make(map[string]minrkt.Exp)
	defineFuncParser := make(map[string]minrkt.Exp)
	temporaryMap := make(map[string]minrkt.Exp)
	DefineStorage := minrkt.Hashmap_minrkt{DefineNameMap: defineNameMap, DefineFuncParser: defineFuncParser, TemporaryMap: temporaryMap}
	stack := minrkt.Stack{StackHM: []minrkt.Hashmap_minrkt{DefineStorage}}
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		tokens, err := minrkt.Tokenizer(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		exp := minrkt.Parser(tokens)
		if exp == nil {
			continue
		}

		minrkt.PrintResult(minrkt.Evaluator(exp, stack))

	}
}
