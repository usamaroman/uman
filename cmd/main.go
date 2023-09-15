package main

import (
	"log"
	"os"

	"uman/repl"
)

func main() {
	log.Println("UMAN")

	args := os.Args
	switch len(args) {
	case 1:
		log.Println("run REPL")
	case 2:
		repl.ReadFile(args[1])
	default:
		log.Fatal("wrong command")
	}

}
