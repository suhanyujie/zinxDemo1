package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT 标识符
	IDENT = "IDENT"
	INT   = "int"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var Keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"int": INT,
}