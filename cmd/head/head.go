package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// readStream reads from a stream (type io.Reader) until EOF and returns the data read as a byte slice.
func readStream(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Head reads from a stream until EOF and returns the first n bytes of the stream.
func main() {
	var data []byte
	var lc = flag.Int("c", 10, "number of bytes to print")
	var err error

	data, err = readStream(os.Stdin)
	if err != nil {
		panic(err)
	}

	for i := 0; i < *lc; i++ {
		fmt.Print(string(data[i]))
	}
}
