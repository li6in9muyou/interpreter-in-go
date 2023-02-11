package ast

import (
	"fmt"
	"strings"
)

type CallExpression struct {
	Function   IExpr
	Parameters []IExpr
}

func (c CallExpression) TokenLiteral() string {
	return c.Function.TokenLiteral()
}

func (c CallExpression) String() string {
	var params []string
	for _, p := range c.Parameters {
		params = append(params, p.TokenLiteral())
	}

	return fmt.Sprintf(
		"%s(%s)",
		c.TokenLiteral(),
		strings.Join(params, ", "),
	)
}

func (c CallExpression) expression() {
	//TODO implement me
	panic("implement me")
}
