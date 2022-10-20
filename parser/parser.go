package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
	statements   []ast.Statement
	errors       []error
}

func (parser *Parser) Errors() []error {
	return parser.errors
}

func (parser *Parser) addError(err error) {
	if err != nil {
		parser.errors = append(parser.errors, err)
	}
}

func New(lexer *lexer.Lexer) *Parser {
	var _ error
	parser := Parser{
		lexer: lexer,
	}
	parser.currentToken, _ = parser.lexer.NextToken()
	parser.nextToken, _ = parser.lexer.NextToken()
	return &parser
}

func (parser *Parser) ParseProgram() (*ast.Program, error) {
	for !parser.currentTokenIs(token.EOF) {
		switch parser.currentToken.Class {
		case token.LET:
			{
				stmt, err := parser.tryLetStatement()
				parser.statements = append(parser.statements, &stmt)
				if err != nil {
					return &ast.Program{Statements: parser.statements},
						fmt.Errorf("%v", err)
				}
			}
		case token.RETURN:
			{
				stmt, err := parser.tryReturnStatement()
				parser.statements = append(parser.statements, &stmt)
				if err != nil {
					return &ast.Program{Statements: parser.statements},
						fmt.Errorf("%v", err)
				}
			}
		default:
			{
				parser.addError(fmt.Errorf(
					"parser error: not implemented %+v",
					parser.currentToken,
				))
				return &ast.Program{Statements: parser.statements}, nil
			}
		}
	}
	return &ast.Program{Statements: parser.statements}, nil
}

func (parser *Parser) tryLetStatement() (ast.LetStatement, error) {
	var err error
	stmt := ast.LetStatement{}

	stmt.Token = parser.currentToken
	parser.eatToken()

	stmt.Name, err = parser.tryIdentExpr()
	if err != nil {
		return stmt, err
	}
	parser.eatToken()

	err = parser.tryAssignOp()
	if err != nil {
		return stmt, err
	}
	parser.eatToken()

	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.eatToken()
	}
	parser.eatToken()

	return stmt, nil
}

func (parser *Parser) tryIdentExpr() (ast.Identifier, error) {
	err := parser.errorCurrentTokenMismatch(token.IDENT)
	if err != nil {
		parser.addError(err)
		return ast.Identifier{}, err
	}
	return ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}, nil
}

func errorTokenMismatch(actual token.Token, expected token.Class) error {
	if actual.Class != expected {
		return fmt.Errorf("expected class %v, got %v", expected, actual)
	}
	return nil
}

func (parser *Parser) errorCurrentTokenMismatch(expected token.Class) error {
	return errorTokenMismatch(parser.currentToken, expected)
}

func (parser *Parser) currentTokenIs(expected token.Class) bool {
	return nil == errorTokenMismatch(parser.currentToken, expected)
}

func (parser *Parser) tryAssignOp() error {
	return parser.errorCurrentTokenMismatch(token.ASSIGN)
}

func (parser *Parser) eatToken() {
	parser.currentToken = parser.nextToken
	parser.nextToken, _ = parser.lexer.NextToken()
}

func (parser *Parser) tryReturnStatement() (ast.ReturnStatement, error) {
	stmt := ast.ReturnStatement{}

	stmt.Token = parser.currentToken
	parser.eatToken()

	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.eatToken()
	}
	parser.eatToken()

	return stmt, nil
}
