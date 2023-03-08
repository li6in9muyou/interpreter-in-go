package parser

import (
	"interpreter/ast"
	"interpreter/token"
)

func (parser *Parser) tryBlockStatement() ast.IExpr {
	parser.eatToken()
	block := ast.BlockStatement{
		OpeningBracket: token.Token{
			Literal: "{",
			Class:   token.LBRACE,
		},
	}
	for !parser.currentTokenIs(token.RBRACE) && !parser.currentTokenIs(token.EOF) {
		stmt, _ := parser.tryStatement()
		block.Statements = append(block.Statements, stmt)
		parser.eatToken()
	}
	parser.eatToken()
	return block
}
