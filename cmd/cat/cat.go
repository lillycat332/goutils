package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readStream(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var data []byte
	var err error

	data, err = readStream(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))
}
