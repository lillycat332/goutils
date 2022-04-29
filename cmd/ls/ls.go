package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ls lists files in a folder.
// ls does NOT hide files starting with a dot, as this was never an intended behaviour in Unix.
func main() {
	var dir string

	if len(os.Args) <= 1 {
		dir = "./"
	} else {
		dir = os.Args[1]
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
