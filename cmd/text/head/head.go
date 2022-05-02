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
func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] [file]...:\n", os.Args[0])

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "	-%s: %v (default: %s)\n", f.Name, f.Usage, f.DefValue)
		})
	}
}

// Head reads from a stream until EOF and returns the first n bytes of the stream.
// Usage: head -b [10] -f [file]...
func main() {
	var data []byte
	var lc = flag.Int("b", 10, "number of bytes to print")
	var file = flag.String("f", "", "file to read")
	var err error

	flag.Parse()
	if *file == "" {
		data, err = readStream(os.Stdin)
		if err != nil {
			fmt.Printf("head: %s", err)
			os.Exit(1)
		}
	} else {
		data, err = ioutil.ReadFile(*file)
		if err != nil {
			fmt.Printf("head: %s", err)
			os.Exit(1)
		}
	}

	if err != nil {
		fmt.Printf("head: %s", err)
		os.Exit(1)
	}

	for i := 0; i < *lc; i++ {
		fmt.Print(string(data[i]))
	}
}
