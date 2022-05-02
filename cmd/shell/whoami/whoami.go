package main

import "os/user"

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	print(u.Username)
}
