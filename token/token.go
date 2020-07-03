package token

// Package token declares the tokens of the language. An input file is broken down into
// a stream of tokens by the lexer and passed to the parser.

// TokenType represents a token. Real languages would use a more compact representation
// for the sake of performance. But, to avoid a lot of boilerplate code to convert to a
// string for printing, we simple use a string.
type TokenType string

// All tokens in language.
const (
	// Special token used to signal errors.
	ILLEGAL = "ILLEGAL"

	// Special token used for end of file.
	EOF = "EOF"

	// Identifier: a reference to a variable binding or function.
	IDENT = "IDENT"

	// Integer literals
	INT = "INT"

	// String literals
	STRING = "STRING"

	// Operators
	ASSIGN   = "ASSIGN"
	BIND     = ":="
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

type Token struct {
	Type     TokenType
	Literal  string
	Location Location
}

type Location struct {
	Line   int
	Column int
}


// Creates a Token from a single character.
func NewTokenFromCharacter(t TokenType, ch byte, loc Location) Token {
	return Token{
		Type:     t,
		Literal:  string(ch),
		Location: loc,
	}
}

// Creates a token from a literal string.
func NewTokenFromLiteral(t TokenType, lit string, loc Location) Token {
	newLoc := loc
	newLoc.Column -= len(lit) - 1
	return Token{
		Type:     t,
		Literal:  lit,
		Location: newLoc,
	}
}

// LookupIdent looks up an identifier and returns the identifier's token, either a
// keyword or binding.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}

// Maps keywords to their tokens.
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
