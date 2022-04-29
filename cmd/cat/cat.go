package main

import (
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

// cat con*CAT*enates files and writes the result to stdout.
// usage: cat [file]...
func main() {
	var data []byte
	var err error

	data, err = readStream(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))
}
