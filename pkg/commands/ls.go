package commands

import (
	"fmt"
	"io/ioutil"
)

func Ls(args []string) {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
