package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ls lists files in a folder.
// ls does NOT hide files starting with a dot, as this was never an intended behaviour in Unix.
// usage: ls [folder]
func main() {
	var dir string

	if len(os.Args) <= 1 {
		dir = "./"
	} else {
		dir = os.Args[1]
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("ls: %s", err)
		os.Exit(1)
	}

	// Iterate&print over files, print them with "/" at the end to indicate a folder.
	for _, file := range files {
		if file.IsDir() {
			fmt.Print(file.Name(), "/", "\n")
		} else {
			fmt.Print(file.Name(), "\n")
		}
	}
}
