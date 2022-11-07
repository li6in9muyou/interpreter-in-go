package ast

import (
	"bytes"
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
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (b BlockStatement) expression() {
	//TODO implement me
	panic("implement me")
}
