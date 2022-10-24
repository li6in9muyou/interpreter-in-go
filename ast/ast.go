package ast

import (
	"bytes"
	"interpreter/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) String() string {
	return i.Value
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

func (statement *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(statement.TokenLiteral() + " ")
	out.WriteString(statement.Name.String())
	out.WriteString(" = ")
	if statement.Value != nil {
		out.WriteString(statement.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value IExpr
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statement() {}

func (r *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " ")
	if r.Value != nil {
		out.WriteString(r.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression IExpr
}

func (e ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e ExpressionStatement) statement() {}

func (e ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (i IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i IntegerLiteral) String() string {
	return i.TokenLiteral()
}

func (i IntegerLiteral) expression() {
}
