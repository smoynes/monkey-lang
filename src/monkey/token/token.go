package token

// Type of a token
type TokenType string

// All tokens in language.
const (
	// Special token used to signal errors.
	ILLEGAL = "ILLEGAL"

	// Special token used for end of file.
	EOF = "EOF"

	// Identifier
	IDENT = "IDENT"

	// Integer literals
	INT = "INT"

	// String literals
	STRING = "STRING"

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
	TILDE    = "~"

	// Comparison operators
	LT  = "<"
	GT  = ">"
	EQ  = "=="
	NEQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	// Block comments
	COMMENT = "COMMENT"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}

}
