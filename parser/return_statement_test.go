package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func Test_tryReturnStatements(t *testing.T) {
	input := `
return 5;
return y+y;
return add(1,2)+product(3,4);
`
	l := lexer.New(input)
	p := New(&l)
	program, _ := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}
