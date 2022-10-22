package ast

import (
	"interpreter/token"
	"testing"
)

func Test_ASTTreeStringer(t *testing.T) {
	expected := "let myVar = anotherVar;let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Class: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Class: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Class: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
			&LetStatement{
				Token: token.Token{Class: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Class: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Class: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
