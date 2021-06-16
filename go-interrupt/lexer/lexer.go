package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
