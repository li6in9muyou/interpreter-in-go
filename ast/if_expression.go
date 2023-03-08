package ast

import (
	"fmt"
	"interpreter/token"
)

type IfExpression struct {
	Token     token.Token
	Predicate IExpr
	Then      *BlockStatement
	Else      *BlockStatement
}

func (i IfExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i IfExpression) String() string {
	return fmt.Sprintf("%s%s %selse %s)", i.Token.Literal, i.Predicate, i.Then, i.Else)
}

func (i IfExpression) expression() {}
