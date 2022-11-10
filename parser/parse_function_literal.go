package parser

import (
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryFunctionLiteral() ast.IExpr {
	t := parser.currentToken
	parser.eatToken()
	parameters := parser.tryFunctionParameters()
	b := (parser.tryBlockStatement()).(ast.BlockStatement)
	return ast.FunctionLiteral{
		Token:      t,
		Parameters: parameters,
		Body:       b,
	}
}

func (parser *Parser) tryFunctionParameters() []ast.Identifier {
	var ans []ast.Identifier
	parser.eatToken()
	for !parser.currentTokenIs(token.RPAREN) {
		ans = append(ans, *parser.tryIdentifierExpr().(*ast.Identifier))
		if parser.currentTokenIs(token.COMMA) {
			parser.eatToken()
		}
	}
	parser.eatToken()
	return ans
}
