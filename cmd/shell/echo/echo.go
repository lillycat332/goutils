package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo writes its arguments to stdout.
// Echo does not have any flags.
// usage: echo [string]
func main() {
	if len(os.Args) < 1 {
		os.Exit(1)
	} else {
		fmt.Print(strings.Join(os.Args[1:], " "))
	}
}
