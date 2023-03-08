package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
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
	dictPrecedence       map[token.Class]int
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
	parser.dictPrecedence = map[token.Class]int{
		token.EQUAL:    EQUALS,
		token.UNEQUAL:  EQUALS,
		token.LT:       LESSGREATER,
		token.GT:       LESSGREATER,
		token.PLUS:     SUM,
		token.MINUS:    SUM,
		token.SLASH:    PRODUCT,
		token.ASTERISK: PRODUCT,
		token.LPAREN:   CALL,
	}
	parser.currentToken, _ = parser.lexer.NextToken()
	parser.nextToken, _ = parser.lexer.NextToken()

	parser.prefixParseFunctions = make(map[token.Class]prefixParseFunction)
	parser.addPrefixFn(token.IDENT, parser.tryIdentifierExpr)
	parser.addPrefixFn(token.INT, parser.tryIntegerLiteralExpr)
	parser.addPrefixFn(token.BANG, parser.tryPrefixExpr)
	parser.addPrefixFn(token.MINUS, parser.tryPrefixExpr)
	parser.addPrefixFn(token.FALSE, parser.tryBooleanLiteralExpr)
	parser.addPrefixFn(token.TRUE, parser.tryBooleanLiteralExpr)
	parser.addPrefixFn(token.LPAREN, parser.tryGroupedExpr)
	parser.addPrefixFn(token.IF, parser.tryIfExpr)
	parser.addPrefixFn(token.LBRACE, parser.tryBlockStatement)
	parser.addPrefixFn(token.FUNCTION, parser.tryFunctionLiteral)

	parser.infixParseFunctions = make(map[token.Class]infixParseFunction)
	parser.addInfixFn(token.PLUS, parser.tryInfixExpr)
	parser.addInfixFn(token.PLUS, parser.tryInfixExpr)
	parser.addInfixFn(token.MINUS, parser.tryInfixExpr)
	parser.addInfixFn(token.SLASH, parser.tryInfixExpr)
	parser.addInfixFn(token.ASTERISK, parser.tryInfixExpr)
	parser.addInfixFn(token.EQUAL, parser.tryInfixExpr)
	parser.addInfixFn(token.UNEQUAL, parser.tryInfixExpr)
	parser.addInfixFn(token.LT, parser.tryInfixExpr)
	parser.addInfixFn(token.GT, parser.tryInfixExpr)
	parser.addInfixFn(token.LPAREN, parser.tryCallExpr)
	return &parser
}

func (parser *Parser) ParseProgram() (*ast.Program, error) {
	for !parser.currentTokenIs(token.EOF) {
		stmt, err := parser.tryStatement()
		if parser.currentTokenIs(token.SEMICOLON) {
			parser.eatToken()
		}
		parser.statements = append(parser.statements, stmt)
		if err != nil {
			return &ast.Program{Statements: parser.statements},
				fmt.Errorf("%v", err)
		}
	}
	return &ast.Program{Statements: parser.statements}, nil
}

func (parser *Parser) tryStatement() (ast.Statement, error) {
	switch parser.currentToken.Class {
	case token.LBRACE:
		{
			stmt := (parser.tryBlockStatement()).(ast.Statement)
			return stmt, nil
		}
	case token.LET:
		{
			stmt, err := parser.tryLetStatement()
			return &stmt, err
		}
	case token.RETURN:
		{
			stmt, err := parser.tryReturnStatement()
			return &stmt, err
		}
	default:
		{
			stmt, ok := parser.tryExpressionStatement()
			if !ok {
				return nil, fmt.Errorf("parse statement failed")
			}
			return &stmt, nil
		}
	}
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

func (parser *Parser) nextTokenIs(expected token.Class) bool {
	return nil == errorTokenMismatch(parser.nextToken, expected)
}

func (parser *Parser) tryAssignOp() error {
	return parser.errorCurrentTokenMismatch(token.ASSIGN)
}

func (parser *Parser) eatToken() {
	parser.currentToken = parser.nextToken
	parser.nextToken, _ = parser.lexer.NextToken()
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

func (parser *Parser) tryExpressionStatement() (ast.ExpressionStatement, bool) {
	stmt := ast.ExpressionStatement{Token: parser.currentToken}
	stmt.Expression = parser.tryExpression(LOWEST)
	if stmt.Expression == nil {
		return stmt, false
	}
	return stmt, true
}

func (parser *Parser) tryExpression(precedence int) ast.IExpr {
	prefix, ok := parser.prefixParseFunctions[parser.currentToken.Class]
	if !ok {
		parser.addError(fmt.Errorf(
			"no prefix parse function for %T%+v",
			parser.currentToken, parser.currentToken,
		))
		return nil
	}
	leftExpr := prefix()
	for !parser.nextTokenIs(token.SEMICOLON) && precedence < parser.currentTokenPrecedence() {
		infix := parser.infixParseFunctions[parser.currentToken.Class]
		if infix == nil {
			return leftExpr
		}

		leftExpr = infix(leftExpr)
	}
	return leftExpr
}

func (parser *Parser) tryPrefixExpr() ast.IExpr {
	op := parser.currentToken
	parser.eatToken()
	right := parser.tryExpression(PREFIX)

	return &ast.PrefixExpression{
		Operator: op,
		Right:    right,
	}
}

func (parser *Parser) tryInfixExpr(left ast.IExpr) ast.IExpr {
	expr := &ast.InfixExpression{
		Operator: parser.currentToken,
		Left:     left,
	}

	precedence := parser.currentTokenPrecedence()
	parser.eatToken()

	expr.Right = parser.tryExpression(precedence)
	return expr
}

func (parser *Parser) currentTokenPrecedence() int {
	if p, ok := parser.dictPrecedence[parser.currentToken.Class]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) nextTokenPrecedence() int {
	if p, ok := parser.dictPrecedence[parser.nextToken.Class]; ok {
		return p
	}
	return LOWEST
}
