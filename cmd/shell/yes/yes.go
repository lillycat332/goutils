package main

import (
	"os"
	"strings"
)

// Set the buffer size as a constant.
const bufsize = 1024 * 1024

func main() {
	y := "y\n"

	// If any arguments are passed, use them as the output, otherwise use y.
	if len(os.Args) > 1 {
		y = strings.Join(os.Args[1:], " ") + "\n"
	}

	buffer := strings.Repeat(y, bufsize/len(y))

	// yyyyyyyyyyyyyyyyyyyyyyyyyyyyyy
	for {
		os.Stdout.Write([]byte(buffer))
	}
}
