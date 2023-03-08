package parser

import "interpreter/ast"

func (parser *Parser) tryReturnStatement() (ast.ReturnStatement, error) {
	stmt := ast.ReturnStatement{}

	stmt.Token = parser.currentToken
	parser.eatToken()

	stmt.Value = parser.tryExpression(LOWEST)
	parser.eatToken()

	return stmt, nil
}
