package lexer

import (
	"go-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	// Create a slice of test structs
	tests := []struct {
		expectedType    token.TokenType
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
		{token.EOF, ""},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d]  tokentype wrong. expected %q, got %q", i, testToken.expectedType, token.Type)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] tokenliteral wrong. expected %q, got %q", i, testToken.expectedLiteral, token.Literal)
		}
	}
}
