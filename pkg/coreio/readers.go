package coreio

import (
	"io"
	"io/ioutil"
)

// readStream reads from a stream (type io.Reader) until EOF and returns the data read as a byte slice.
func readStream(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return data, nil
}
