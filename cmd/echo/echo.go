package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		os.Exit(1)
	} else {
		fmt.Print(strings.Join(os.Args[1:], " "))
	}
}
