package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"uman/repl"
)

func main() {
	log.Println("Uman started")

	args := os.Args
	switch len(args) {
	case 1:
		fmt.Println("Welcome to the Uman progrmming language repl!")
		repl.Start(os.Stdin)
	case 2:
		file, err := os.Open(args[1])
		if err != nil {
			log.Fatal(err)
		}

		text := bufio.NewReader(file)
		repl.Start(text)
	default:
		log.Fatal("wrong argument")
	}

}
