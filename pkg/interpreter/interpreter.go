package interpreter

import (
	"dan/go-shell/pkg/commands"
	"dan/go-shell/pkg/parser"
	"fmt"
	"os"
)

var commandMap = map[string]func([]string){
	"echo": echo,
	"cd":   commands.Cd,
	"ls":   commands.Ls,
	"exit": exit,
}

func Execute(command parser.Command) {
	if function, exists := commandMap[command.Name]; exists {
		function(command.Args)
	} else {
		fmt.Printf("Unknown command: %s\n", command.Name)
	}
}

func echo(args []string) {
	fmt.Println(args)
}

func exit(args []string) {
	os.Exit(0)
}
