package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statement()
}

type IExpr interface {
	Node
	expression()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expression() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value IExpr
}

func (statement *LetStatement) TokenLiteral() string {
	return statement.Token.Literal
}

func (statement *LetStatement) statement() {}

type ReturnStatement struct {
	Token token.Token
	Value IExpr
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statement() {}
