package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey language!\n", user.Username)
	fmt.Printf("Type in some commmands\n")
	repl.Start(os.Stdin, os.Stdout)

}
