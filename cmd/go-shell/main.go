package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dan/go-shell/pkg/interpreter"
	"dan/go-shell/pkg/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if input == "exit" {
			break
		}

		command := parser.Parse(input)
		interpreter.Execute(command)
	}
}
