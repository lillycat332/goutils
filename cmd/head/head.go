package main

import (
	"flag"
	"fmt"
	"os"
)

// Head reads from a stream until EOF and returns the first n bytes of the stream.
func main() {
	var data []byte
	var lc = flag.Int("b", 10, "number of bytes to print")
	var err error
	flag.Parse()

	data, err = readStream(os.Stdin)
	if err != nil {
		panic(err)
	}

	for i := 0; i < *lc; i++ {
		fmt.Print(string(data[i]))
	}
}
