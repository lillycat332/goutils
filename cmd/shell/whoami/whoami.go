package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Printf("whoami: %s", err)
		os.Exit(1)
	}
	print(u.Username)
}
