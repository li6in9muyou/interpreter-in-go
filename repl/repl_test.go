package repl

import (
	"fmt"
	"github.com/repeale/fp-go"
	"strings"
	"testing"
)

type Output []string

func (output *Output) Write(p []byte) (n int, err error) {
	*output = append(*output, string(p))
	return len(p), nil
}

func TestItShouldPrintPromptAndBye(t *testing.T) {
	input := strings.NewReader("")
	var output Output
	Start(input, &output)

	prompt := fp.Some(
		func(line string) bool { return line == PROMPT },
	)(output)
	bye := fp.Some(
		func(line string) bool { return strings.Contains(strings.ToLower(line), "bye") },
	)(output)

	if !bye || !prompt {
		t.Log(output)
		if !bye {
			t.Fatalf("bye is not shown")
		}
		if !prompt {
			t.Fatalf("prompt is not shown")
		}
	}
}

func TestItShouldReportIllegalLines(t *testing.T) {
	const bad = "\u0007"
	input := strings.NewReader(fmt.Sprintf("%v\n%v\n", bad, QUIT))
	var output Output
	Start(input, &output)

	if !fp.Some(
		func(line string) bool {
			return strings.Contains(strings.ToLower(line), "illegal") &&
				strings.Contains(strings.ToLower(line), bad)
		},
	)(output) {
		t.Log(output)
		t.Fatalf("illagal input is not reported")
	}
}
