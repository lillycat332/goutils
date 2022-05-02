package main

import (
	"os"
	"strings"
)

const bufsize = 1024 * 1024

func main() {
	y := "y\n"

	if len(os.Args) > 1 {
		y = strings.Join(os.Args[1:], " ") + "\n"
	}

	buffer := strings.Repeat(y, bufsize/len(y))

	for {
		os.Stdout.Write([]byte(buffer))
	}
}
