package parser

import (
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryFunctionLiteral() ast.IExpr {
	t := parser.currentToken
	parser.eatToken() // `fun` keyword
	name, _ := parser.tryIdentExpr()
	parser.eatToken() // identifier
	parameters := parser.tryFunctionParameters()
	b := (parser.tryBlockStatement()).(ast.BlockStatement)
	return ast.FunctionLiteral{
		Token:        t,
		FunctionName: name.Token,
		Parameters:   parameters,
		Body:         b,
	}
}

func (parser *Parser) tryFunctionParameters() []ast.Identifier {
	var ans []ast.Identifier
	parser.eatToken() // left parenthesis
	for !parser.currentTokenIs(token.RPAREN) {
		ans = append(ans, *parser.tryIdentifierExpr().(*ast.Identifier))
		if parser.currentTokenIs(token.COMMA) {
			parser.eatToken()
		}
	}
	parser.eatToken()
	return ans
}
