package ast

import (
	"bytes"
	"interpreter/token"
	"strings"
)

type FunctionLiteral struct {
	Token        token.Token
	FunctionName token.Token
	Parameters   []Identifier
	Body         BlockStatement
}

func (f FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f FunctionLiteral) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(f.TokenLiteral())
	out.WriteString(f.FunctionName.Literal)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}

func (f FunctionLiteral) expression() {
	//TODO implement me
	panic("implement me")
}
