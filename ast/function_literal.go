package ast

import (
	"fmt"
	"interpreter/token"
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
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	text := fmt.Sprintf(
		"fun %s(%s) {\n%s}\n",
		f.TokenLiteral(),
		f.FunctionName.Literal,
		f.Body.String(),
	)
	return text
}

func (f FunctionLiteral) expression() {
	//TODO implement me
	panic("implement me")
}
