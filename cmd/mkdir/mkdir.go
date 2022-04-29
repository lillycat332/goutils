package main

import (
	"fmt"
	"os"
)

// Mkdir creates a new directory with the specified name.
// TODO: Add support for permision bits.
func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: mkdir <dir>")
		os.Exit(1)
	} else {
		err := os.Mkdir(os.Args[1], 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
