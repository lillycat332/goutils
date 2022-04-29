package main

import (
	"fmt"
	"os"
	"strings"
)

// echo writes its arguments to stdout.
func main() {
	if len(os.Args) < 1 {
		os.Exit(1)
	} else {
		fmt.Print(strings.Join(os.Args[1:], " "))
	}
}
