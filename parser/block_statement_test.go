package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func Test_parseBlockExpression(t *testing.T) {
	input := `{ let x=6;let y=(x+8)*x;let x=x+y;let y=2*x;return x<y; }`
	l := lexer.New(input)
	p := New(&l)
	program, _ := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	block, ok := (program.Statements[0]).(ast.BlockStatement)
	fmt.Println(block)
	if !ok {
		t.Fatalf("expecting ast.BlockStatement, got %T", block)
	}
	if len(block.Statements) != 5 {
		t.Fatalf("expecting 5 statements, got %v", len(block.Statements))
	}
}
