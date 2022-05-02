package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	err := syscall.Fsync(0)
	if err != nil {
		fmt.Printf("ls: %s", err)
		os.Exit(1)
	}
}
