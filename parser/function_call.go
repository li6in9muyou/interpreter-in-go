package parser

import (
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryCallExpr(callee ast.IExpr) ast.IExpr {
	var parameters []ast.IExpr
	parser.eatToken()

	if parser.nextTokenIs(token.RPAREN) {
		parser.eatToken()
		return ast.CallExpression{Function: callee, Parameters: parameters}
	}

	parameters = append(parameters, parser.tryExpression(LOWEST))
	for parser.currentTokenIs(token.COMMA) {
		parser.eatToken()
		parameters = append(parameters, parser.tryExpression(LOWEST))
	}

	parser.eatToken()

	return ast.CallExpression{Function: callee, Parameters: parameters}
}
