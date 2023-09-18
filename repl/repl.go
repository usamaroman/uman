package repl

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"uman/lexer"
	"uman/token"
)

var ErrWrongExtension = errors.New("wrong file extension")

func Run() {
	const prompt = ">> "
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		l := lexer.New(text)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Println(prompt, tok)
		}

	}

}

func ReadFile(filename string) {
	err := readFileExtension(filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for {
		if ok := scanner.Scan(); !ok {
			os.Exit(1)
		}

		input := scanner.Text()
		l := lexer.New(input)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}

func readFileExtension(filename string) error {
	split := strings.Split(filename, ".")
	switch split[len(split)-1] {
	case "um":
		return nil
	default:
		return ErrWrongExtension
	}
}
