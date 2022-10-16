package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
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

		for lex := lexer.New(line); ; {
			t, err := lex.NextToken()

			if t.Class == token.EOF {
				show("Done.\n")
				break
			}
			if err != nil {
				show("Error: %v\n", err)
				break
			}

			show("%+v\n", t)
		}
	}
	_, _ = fmt.Fprintln(out, "Bye!")
}
