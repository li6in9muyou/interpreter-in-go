package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"Hello %s! This is the Monkey programming language!\n",
		current.Username,
	)
	repl.Start(os.Stdin, os.Stdout)
}
