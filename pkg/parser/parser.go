package parser

import (
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func Parse(input string) Command {
	parts := strings.Fields(input)

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}
}
