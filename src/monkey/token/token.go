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
	WHILE    = "WHILE"

	// Block comments
	COMMENT = "COMMENT"
)

func NewTokenFromCharacter(t TokenType, ch byte, loc Location) Token {
	return Token{
		Type:     t,
		Literal:  string(ch),
		Location: loc,
	}
}

func NewTokenFromLiteral(t TokenType, lit string, loc Location) Token {
	newLoc := loc
	newLoc.Column -= len(lit) - 1
	return Token{
		Type:     t,
		Literal:  lit,
		Location: newLoc,
	}
}

type Token struct {
	Type     TokenType
	Literal  string
	Location Location
}

type Location struct {
	Line   int
	Column int
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"while":  WHILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}
