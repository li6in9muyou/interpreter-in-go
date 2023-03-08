package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryGroupedExpr() ast.IExpr {
	parser.eatToken()
	expr := parser.tryExpression(LOWEST)
	if parser.currentTokenIs(token.RPAREN) {
		parser.eatToken()
		return expr
	}

	parser.addError(fmt.Errorf(
		"there is no right parenthesis after %s", expr,
	))
	return nil
}
