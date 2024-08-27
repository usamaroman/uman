package main

import (
	"log"
	"os"

	"github.com/usamaroman/uman/repl"
)

func main() {
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
