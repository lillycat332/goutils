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
		os.Exit(0)
	} else {
		// Echo the arguments joined together
		fmt.Print(strings.Join(os.Args[1:], " "))
	}
}
