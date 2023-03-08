package parser

import (
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryIfExpr() ast.IExpr {
	expr := ast.IfExpression{Token: parser.currentToken}

	parser.eatToken()

	if parser.currentTokenIs(token.LPAREN) {
		parser.eatToken()
		expr.Predicate = parser.tryExpression(LOWEST)
	}
	parser.eatToken()

	var b = (parser.tryBlockStatement()).(ast.BlockStatement)
	expr.Then = &b
	return expr
}
