package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	err := syscall.Fsync(0)
	if err != nil {
		fmt.Printf("sync: %s", err)
		os.Exit(1)
	}
}
