package lexer

import (
	"fmt"
	"github.com/suhanyujie/go-interrupt/token"
	"testing"
)

type TokenTestT1 struct {
	expectedType    token.TokenType
	expectedLiteral string
}

/// 我们可以识别单个字符的 token，那如何拓展到多字符组成的 token 呢？
func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	testCases := []TokenTestT1{
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
	for i, tt := range testCases {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("[test] tokenType wrong... index: %d\n", i)
			return
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("[test] token literal wrong... index: %d\n", i)
			return
		}
	}
}

func TestNextTokenV2(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
!-/*5;
5<10 > 9
if (5 < 10) {
	return true
} else {
	return false
}
`
	_ = []TokenTestT1{
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
		{token.FUNCTION, "fn"},
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
		{token.EOF, ""},
	}
	l := New(input)
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}
}

func TestIsLetter1(t *testing.T) {
	if !isLetter('a') {
		t.Errorf("isLetter error. 1")
		return
	}
	if isLetter('1') {
		t.Errorf("isLetter error. 2")
		return
	}
}
