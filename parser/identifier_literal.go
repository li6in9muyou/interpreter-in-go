package parser

import (
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryIdentifierExpr() ast.IExpr {
	identifier := parser.currentToken
	parser.eatToken()
	return &ast.Identifier{
		Token: identifier,
		Value: identifier.Literal,
	}
}

func (parser *Parser) tryIdentExpr() (ast.Identifier, error) {
	err := parser.errorCurrentTokenMismatch(token.IDENT)
	if err != nil {
		parser.addError(err)
		return ast.Identifier{}, err
	}
	return ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}, nil
}
