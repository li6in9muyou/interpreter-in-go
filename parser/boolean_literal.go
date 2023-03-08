package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryBooleanLiteralExpr() ast.IExpr {
	t := parser.currentToken
	parser.eatToken()

	switch t.Class {
	case token.FALSE:
		{
			return &ast.BooleanLiteral{Token: t, Value: false}
		}
	case token.TRUE:
		{
			return &ast.BooleanLiteral{Token: t, Value: true}
		}
	}

	parser.addError(fmt.Errorf("unknown boolean literal %s", t.Literal))
	return nil
}
