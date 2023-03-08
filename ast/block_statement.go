package ast

import (
	"bytes"
	"fmt"
	"interpreter/token"
)

type BlockStatement struct {
	OpeningBracket token.Token
	Statements     []Statement
}

func (b BlockStatement) statement() {
	//TODO implement me
	panic("implement me")
}

func (b BlockStatement) TokenLiteral() string {
	return b.OpeningBracket.Literal
}

func (b BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for _, s := range b.Statements {
		out.WriteString(fmt.Sprintf("  %v\n", s))
	}
	out.WriteString("}\n")
	return out.String()
}

func (b BlockStatement) expression() {
	//TODO implement me
	panic("implement me")
}
