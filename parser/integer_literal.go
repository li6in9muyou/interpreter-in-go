package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/token"
	"strconv"
)

func (parser *Parser) tryIntegerLiteralExpr() ast.IExpr {
	t := parser.currentToken
	parser.eatToken()

	number, err := strconv.Atoi(t.Literal)
	if err != nil {
		parser.addError(fmt.Errorf(
			"%s can not be parsed as base 10 integer", t.Literal,
		))
		return &ast.IntegerLiteral{
			Token: token.New(token.ILLEGAL, t.Literal),
			Value: 0,
		}
	}
	return &ast.IntegerLiteral{
		Token: t,
		Value: number,
	}
}
