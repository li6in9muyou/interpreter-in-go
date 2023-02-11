package parser

import (
	"interpreter/lexer"
	"strings"
	"testing"
)

func Test_tryLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = x;
let y = y + 2;
let y = add(y,x)+y;
`
	l := lexer.New(input)
	p := New(&l)
	program, _ := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	expectedStatementCount := strings.Count(input, ";")
	if len(program.Statements) != expectedStatementCount {
		t.Fatalf(
			"program.Statements does not contain %d statements. got=%d",
			expectedStatementCount,
			len(program.Statements),
		)
	}
}
