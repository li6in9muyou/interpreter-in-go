package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func Test_parseCallExpression(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"
	l := lexer.New(input)
	p := New(&l)
	program, _ := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T",
			stmt.Expression)
	}
	if !testIdentifier(t, &exp.Function, "add") {
		return
	}
	if len(exp.Parameters) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Parameters))
	}
	testLiteralExpression(t, &exp.Parameters[0], 1)
	testInfixExpression(t, &exp.Parameters[1], 2, "*", 3)
	testInfixExpression(t, &exp.Parameters[2], 4, "+", 5)
}
