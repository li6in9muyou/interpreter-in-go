package lexer

import (
	"errors"
	"fmt"
	"interpreter/token"
	"strings"
)

type Lexer struct {
	input    string
	position int
}

var dictAtom = map[string]token.Token{
	"=": token.New(token.ASSIGN, "="),
	"+": token.New(token.PLUS, "+"),
	"(": token.New(token.LPAREN, "("),
	")": token.New(token.RPAREN, ")"),
	"{": token.New(token.LBRACE, "{"),
	"}": token.New(token.RBRACE, "}"),
	",": token.New(token.COMMA, ","),
	"-": token.New(token.MINUS, "-"),
	"!": token.New(token.BANG, "!"),
	"/": token.New(token.SLASH, "/"),
	">": token.New(token.GT, ">"),
	"<": token.New(token.LT, "<"),
	"*": token.New(token.ASTERISK, "*"),
	";": token.New(token.SEMICOLON, ";"),
}

var dictKeyword = map[string]token.Token{
	"let":    token.New(token.LET, "let"),
	"fun":    token.New(token.FUNCTION, "fun"),
	"true":   token.New(token.TRUE, "true"),
	"false":  token.New(token.FALSE, "false"),
	"if":     token.New(token.IF, "if"),
	"else":   token.New(token.ELSE, "else"),
	"return": token.New(token.RETURN, "return"),
	"!=":     token.New(token.UNEQUAL, "!="),
	"==":     token.New(token.EQUAL, "=="),
	"&&":     token.New(token.LOGICAND, "&&"),
	"||":     token.New(token.LOGICOR, "||"),
}

func (lexer *Lexer) eatBlankSpace() {
	if lexer.position >= len(lexer.input) {
		return
	}

	var ch byte
	for lexer.position < len(lexer.input) {
		ch = lexer.input[lexer.position]
		if !(ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r') {
			break
		}
		lexer.position += 1
	}
}

func (lexer *Lexer) eatChar() (char string) {
	lexer.eatBlankSpace()

	char, _ = lexer.peekChar()
	lexer.position += 1
	return
}

func (lexer *Lexer) eatNumber() (char string) {
	lexer.eatBlankSpace()

	ch, err := lexer.peekChar()
	for err == nil && isDigit(ch) {
		char += lexer.eatChar()
		if lexer.position >= len(lexer.input) || lexer.input[lexer.position] == ' ' {
			break
		}
		ch, err = lexer.peekChar()
	}
	return
}

func (lexer *Lexer) eatWord() (char string) {
	lexer.eatBlankSpace()

	ch, err := lexer.peekChar()
	for err == nil && isLetter(ch) {
		char += lexer.eatChar()
		if lexer.position >= len(lexer.input) || lexer.input[lexer.position] == ' ' {
			break
		}
		ch, err = lexer.peekChar()
	}
	return
}

func (lexer *Lexer) peekChar() (string, error) {
	lexer.eatBlankSpace()

	if lexer.position < len(lexer.input)-1 {
		return lexer.input[lexer.position : lexer.position+1], nil
	} else if lexer.position == len(lexer.input)-1 {
		return lexer.input[lexer.position:], nil
	} else {
		return "", errors.New("EOF")
	}
}

func New(input string) Lexer {
	return Lexer{input: input}
}

func isDigit(ch string) bool {
	return '0' <= ch[0] && ch[0] <= '9'
}

func isAtom(s string) bool {
	_, ok := dictAtom[s]
	return ok
}

func isLetter(s string) bool {
	return 'a' <= s[0] && s[0] <= 'z' || 'A' <= s[0] && s[0] <= 'Z' || s[0] == '_'
}

func (lexer *Lexer) NextToken() (token.Token, error) {
	ch, err := lexer.peekChar()

	if err != nil {
		return token.Token{
			Class:   token.EOF,
			Literal: "EOF",
		}, nil
	}

	switch {
	case isAtom(ch) || isPrefixOfMultipleAtoms(ch):
		{
			var word string
			if isPrefixOfMultipleAtoms(ch) {
				word = lexer.eatWhile(isPrefixOfMultipleAtoms)
				if t, err := lexer.tryKeyword(word); err == nil {
					return t, err
				}
			} else {
				word = lexer.eatChar()
			}
			return lexer.tryAtom(word)
		}
	case isLetter(ch):
		{
			var t token.Token
			var err error

			word := lexer.eatWord()
			t, err = lexer.tryKeyword(word)
			if err == nil {
				return t, err
			}

			t, err = lexer.tryIdentifier(word)
			if err == nil {
				return t, err
			}
		}
	case isDigit(ch):
		{
			number := lexer.eatNumber()
			return lexer.tryInteger(number)
		}
	}
	return token.New(token.ILLEGAL, ""),
		fmt.Errorf("illegal token %v at %v", ch, lexer.position)
}

func isPrefixOfMultipleAtoms(ch string) bool {
	return strings.ContainsAny(ch, "&!=|")
}

func (lexer *Lexer) tryInteger(number string) (token.Token, error) {
	return token.New(token.INT, number), nil
}

func (lexer *Lexer) tryAtom(word string) (token.Token, error) {
	t, ok := dictAtom[word]
	if ok {
		return t, nil
	} else {
		return token.Token{}, errors.New("unknown atom")
	}
}

func (lexer *Lexer) tryKeyword(word string) (token.Token, error) {
	if t, ok := dictKeyword[word]; ok {
		return t, nil
	}
	return token.Token{}, errors.New("not a keyword")
}

func (lexer *Lexer) tryIdentifier(word string) (token.Token, error) {
	return token.New(token.IDENT, word), nil
}

func (lexer *Lexer) eatWhile(predicate func(ch string) bool) string {
	ans := ""
	for ch, err := lexer.peekChar(); predicate(ch) && err == nil; {
		ans += lexer.eatChar()
		ch, err = lexer.peekChar()
	}
	return ans
}
