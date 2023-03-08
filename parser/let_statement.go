package parser

import "interpreter/ast"

func (parser *Parser) tryLetStatement() (ast.LetStatement, error) {
	var err error
	var n ast.Identifier
	stmt := ast.LetStatement{Name: &n}

	stmt.Token = parser.currentToken
	parser.eatToken()

	n, err = parser.tryIdentExpr()
	if err != nil {
		return stmt, err
	}
	parser.eatToken()

	err = parser.tryAssignOp()
	if err != nil {
		return stmt, err
	}
	parser.eatToken()

	stmt.Value = parser.tryExpression(LOWEST)

	return stmt, nil
}
