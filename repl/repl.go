package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"uman/lexer"
	"uman/token"
)

func Start(input io.Reader) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(">> ") // prompt

		if scanned := scanner.Scan(); !scanned {
			log.Printf("Repl line is empty")
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
