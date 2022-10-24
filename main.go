package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/chkda/monkey-lang-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to Monkey Programming Language %s!!", user.Username)
	fmt.Printf("\nStart your code\n")
	repl.Start(os.Stdin, os.Stdout)
}
