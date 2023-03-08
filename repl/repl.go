package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/parser"
	"io"
)

const PROMPT = ">> "
const QUIT = ":q"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	show := func(format string, args ...any) {
		_, _ = fmt.Fprintf(out, format, args...)
	}

	show("Type %v to quit\n", QUIT)

	for {
		show(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		if line == ":q" {
			break
		}

		lex := lexer.New(line)
		p := parser.New(&lex)
		program, err := p.ParseProgram()
		if err != nil {
			show("Your fucked up\n")
			for _, msg := range p.Errors() {
				show("\t%s\n", msg)
			}
		} else {
			show("%+v\n", program)
		}

	}
	_, _ = fmt.Fprintln(out, "Bye!")
}
