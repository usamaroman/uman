package main

import (
	"fmt"
	"log"
	"os"

	"uman/repl"
)

func main() {
	fmt.Println("UMAN - язык программирования")

	args := os.Args
	switch len(args) {
	case 1:
		repl.Run()
	case 2:
		repl.ReadFile(args[1])
	default:
		log.Fatal("wrong command")
	}

}
