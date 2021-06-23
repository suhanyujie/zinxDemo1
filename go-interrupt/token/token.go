package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func CheckIdent(text string) TokenType {
	if tokenType, ok := Keywords[text]; ok {
		return tokenType
	}
	return IDENT
}
