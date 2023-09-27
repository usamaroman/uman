package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"uman/evaluator"
	"uman/object"
	"uman/parser"
)

var ErrWrongExtension = errors.New("wrong file extension")

func Run() {
	const prompt = ">> "
	scanner := bufio.NewScanner(os.Stdin)
	out := os.Stdout
	env := object.NewEnvironment()

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.New(line)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
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

	var input string
	out := os.Stdout
	env := object.NewEnvironment()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}
	log.Println(input)

	textScanner := bufio.NewScanner(strings.NewReader(input))
	for textScanner.Scan() {
		log.Println(textScanner.Text())
		p := parser.New(textScanner.Text())
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		log.Println(evaluated)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
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

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
