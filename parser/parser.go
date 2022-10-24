package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
	"strconv"
)

type (
	prefixParseFunction func() ast.IExpr
	infixParseFunction  func(expr ast.IExpr) ast.IExpr
)

type Parser struct {
	lexer                *lexer.Lexer
	currentToken         token.Token
	nextToken            token.Token
	statements           []ast.Statement
	errors               []error
	prefixParseFunctions map[token.Class]prefixParseFunction
	infixParseFunctions  map[token.Class]infixParseFunction
}

func (parser *Parser) addPrefixFn(class token.Class, function prefixParseFunction) {
	parser.prefixParseFunctions[class] = function
}

func (parser *Parser) addInfixFn(class token.Class, function infixParseFunction) {
	parser.infixParseFunctions[class] = function
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

	parser.prefixParseFunctions = make(map[token.Class]prefixParseFunction)
	parser.addPrefixFn(token.IDENT, parser.tryIdentifierExpr)
	parser.addPrefixFn(token.INT, parser.tryIntegerLiteralExpr)
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
		case token.SEMICOLON:
			{
				parser.eatToken()
			}
		default:
			{
				stmt, err := parser.tryExpressionStatement()
				parser.statements = append(parser.statements, &stmt)
				if err != nil {
					return &ast.Program{Statements: parser.statements},
						fmt.Errorf("%v", err)
				}
			}
		}
	}
	return &ast.Program{Statements: parser.statements}, nil
}

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

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

func (parser *Parser) tryExpressionStatement() (ast.ExpressionStatement, error) {
	stmt := ast.ExpressionStatement{Token: parser.currentToken}
	stmt.Expression = parser.tryExpression(LOWEST)
	if stmt.Expression == nil {
		return stmt, fmt.Errorf("what's wrong")
	}
	return stmt, nil
}

func (parser *Parser) tryExpression(precedence int) ast.IExpr {
	prefix, ok := parser.prefixParseFunctions[parser.currentToken.Class]
	if !ok {
		return nil
	}
	leftExpr := prefix()
	return leftExpr
}

func (parser *Parser) tryIdentifierExpr() ast.IExpr {
	identifier := parser.currentToken
	parser.eatToken()
	return &ast.Identifier{
		Token: identifier,
		Value: identifier.Literal,
	}
}

func (parser *Parser) tryIntegerLiteralExpr() ast.IExpr {
	t := parser.currentToken
	parser.eatToken()

	number, err := strconv.Atoi(t.Literal)
	if err != nil {
		parser.addError(fmt.Errorf(
			"%s can not be parsed as base 10 integer", t.Literal,
		))
		return &ast.IntegerLiteral{
			Token: token.New(token.ILLEGAL, t.Literal),
			Value: 0,
		}
	}
	return &ast.IntegerLiteral{
		Token: t,
		Value: number,
	}
}
