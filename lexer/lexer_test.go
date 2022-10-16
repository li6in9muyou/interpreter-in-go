package lexer

import (
	"interpreter/token"
	"testing"
)

func TestLexer_eatChar_peekChar(t *testing.T) {
	input := "         =       +\n(   \n\n  9\n      ){let},;"
	expected := []string{
		"=", "+", "(", "9", ")", "{", "l", "e", "t", "}", ",", ";",
	}

	lex := New(input)

	// peakChar do not advance reading head
	for i := 0; i < 3; i++ {
		peek, _ := lex.peekChar()
		if peek != "=" {
			t.Fatalf("expected=%q, got=%q", "=", peek)
		}
	}

	// eatChar skips blanks
	for i, exp := range expected {
		if actual := lex.eatChar(); actual != exp {
			t.Fatalf("tests[%d] - character wrong. expected=%q, got=%q",
				i, exp, actual)
		}
	}

	// do panic if input is empty
	lex = New("")
	if _, err := lex.peekChar(); err == nil {
		t.Fatalf("did not return EOF")
	}
}

func TestLexer_peekChar_ReadWords(t *testing.T) {
	input := `         let   add RESULT`
	expected := []string{
		"let", "add", "RESULT",
	}

	lex := New(input)
	for i, exp := range expected {
		if actual := lex.eatWord(); actual != exp {
			t.Fatalf("tests[%d] - token wrong. expected=%q, got=%q",
				i, exp, actual)
		}
	}
}

func TestLexer_NextToken_ShouldReadAtom(t *testing.T) {
	input := `         =       +(           ){},;`
	tests := []struct {
		expectedClass   token.Class
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok, _ := lexer.NextToken()
		if tok.Class != tt.expectedClass {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedClass, tok.Class)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLexer_NextToken_ShouldReadKeywords(t *testing.T) {
	input := `let five = 5;
let ten=10;
let add =fun(x, y) {
x +y;};let result = 
add(five, ten);
`
	tests := []struct {
		expectedClass   token.Class
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok, _ := lexer.NextToken()
		if tok.Class != tt.expectedClass {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedClass, tok.Class)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLexer_NextToken_ShouldReadMoreOperators(t *testing.T) {
	input := "!-/*5;\n5 <10 >5;"

	tests := []struct {
		expectedClass   token.Class
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok, _ := lexer.NextToken()
		if tok.Class != tt.expectedClass {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedClass, tok.Class)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLexer_NextToken_ShouldReadMoreKeywords(t *testing.T) {
	input := "if (89<64) {\nreturn true;\n} else {\nreturn false;\n}"

	tests := []struct {
		expectedClass   token.Class
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "89"},
		{token.LT, "<"},
		{token.INT, "64"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok, _ := lexer.NextToken()
		if tok.Class != tt.expectedClass {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedClass, tok.Class)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLexer_NextToken_ShouldReadMultipleCharOperators(t *testing.T) {
	input := "if(89!=64&&true==false||true){\n\n}"

	tests := []struct {
		expectedClass   token.Class
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "89"},
		{token.UNEQUAL, "!="},
		{token.INT, "64"},
		{token.LOGICAND, "&&"},
		{token.TRUE, "true"},
		{token.EQUAL, "=="},
		{token.FALSE, "false"},
		{token.LOGICOR, "||"},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, "EOF"},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok, _ := lexer.NextToken()
		if tok.Class != tt.expectedClass {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedClass, tok.Class)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
