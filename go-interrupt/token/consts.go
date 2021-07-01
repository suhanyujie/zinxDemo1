package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT 标识符
	IDENT = "IDENT"
	INT   = "int"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"
	BANG      = "!"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

// Keywords 关键字标识符到 token 常量的映射
var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"int":    INT,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
