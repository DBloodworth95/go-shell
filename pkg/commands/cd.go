package commands

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	if len(args) == 0 {
		fmt.Println("Expected path argument")
		return
	}

	err := os.Chdir(args[0])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
