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
	var file string

	if os.Args[1] != "" {
		for i := 1; i < len(os.Args); i++ {
			file = os.Args[i]
			tmp, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("cat: %s", err)
				os.Exit(1)
			}
			data = append(data, tmp...)
		}

		if err != nil {
			fmt.Printf("cat: %s", err)
			os.Exit(1)
		}
	} else {
		data, err = readStream(os.Stdin)
		if err != nil {
			fmt.Printf("cat: %s", err)
			os.Exit(1)
		}
	}

	if err != nil {
		fmt.Printf("head: %s", err)
		os.Exit(1)
	}

	fmt.Print(string(data))
}
