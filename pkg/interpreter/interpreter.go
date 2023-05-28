package interpreter

import (
	"dan/go-shell/pkg/commands"
	"dan/go-shell/pkg/parser"
	"fmt"
	"os"
)

func Execute(command parser.Command) {
	switch command.Name {
	case "echo":
		echo(command.Args)
	case "exit":
		exit()
	case "cd":
		commands.Cd(command.Args)
	case "ls":
		commands.Ls(command.Args)
	default:
		fmt.Printf("Unknown command: %s\n", command.Name)
	}
}

func echo(args []string) {
	fmt.Println(args)
}

func exit() {
	os.Exit(0)
}
