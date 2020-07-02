package lexer

import (
	"github.com/smoynes/monkey-lang/token"
)

type Lexer struct {
	input        string
	position     int // current position in input (current char)
	readPosition int // reading position in input (after current token)
	ch           byte

	loc token.Location
}

func New(input string) *Lexer {
	l := &Lexer{input: input, loc: token.Location{0, 0}}

	l.readChar() // read the first char in the input
	return l
}

// Returns the next token in the lexed stream.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Consume whitespace
	l.skipWhitespace()

	// Tokenize current character.
	switch l.ch {

	case '=':
		// If next char is "=" then emit "==" token, otherwise emit assignment
		// token.
		if l.peekChar() == '=' {
			// Consume peeked byte
			l.readChar()
			tok = token.NewTokenFromLiteral(token.EQ, "==", l.loc)
		} else {
			tok = token.NewTokenFromCharacter(token.ASSIGN, l.ch, l.loc)
		}
	case '!':
		// If next char is "=" then emit "!=" token, otherwise emit unary bang operator
		if l.peekChar() == '=' {
			// Consume peeked byte
			l.readChar()
			tok = token.NewTokenFromLiteral(token.NEQ, "!=", l.loc)
		} else {
			tok = token.NewTokenFromCharacter(token.BANG, l.ch, l.loc)
		}
	case '"':
		tok.Type = token.STRING
		tok.Location = l.loc
		tok.Literal = l.readString()

	case '#':
		tok.Type = token.COMMENT
		tok.Location = l.loc
		tok.Literal = l.readEndOfLine()
	case 0:
		// No more tokens.
		tok.Type = token.EOF
		tok.Literal = ""

	// Operators
	case '+':
		tok = token.NewTokenFromCharacter(token.PLUS, l.ch, l.loc)
	case '-':
		tok = token.NewTokenFromCharacter(token.MINUS, l.ch, l.loc)
	case '/':
		tok = token.NewTokenFromCharacter(token.SLASH, l.ch, l.loc)
	case '*':
		tok = token.NewTokenFromCharacter(token.ASTERISK, l.ch, l.loc)
	case '~':
		tok = token.NewTokenFromCharacter(token.TILDE, l.ch, l.loc)
	case '<':
		tok = token.NewTokenFromCharacter(token.LT, l.ch, l.loc)
	case '>':
		tok = token.NewTokenFromCharacter(token.GT, l.ch, l.loc)
	case '(':
		tok = token.NewTokenFromCharacter(token.LPAREN, l.ch, l.loc)
	case ')':
		tok = token.NewTokenFromCharacter(token.RPAREN, l.ch, l.loc)
	case '{':
		tok = token.NewTokenFromCharacter(token.LBRACE, l.ch, l.loc)
	case '}':
		tok = token.NewTokenFromCharacter(token.RBRACE, l.ch, l.loc)
	case '[':
		tok = token.NewTokenFromCharacter(token.LBRACKET, l.ch, l.loc)
	case ']':
		tok = token.NewTokenFromCharacter(token.RBRACKET, l.ch, l.loc)
	case ';':
		tok = token.NewTokenFromCharacter(token.SEMICOLON, l.ch, l.loc)
	case ',':
		tok = token.NewTokenFromCharacter(token.COMMA, l.ch, l.loc)
	case ':':
		tok = token.NewTokenFromCharacter(token.COLON, l.ch, l.loc)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Location = l.loc
			tok.Location.Column -= len(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Location = l.loc
			tok.Location.Column -= len(tok.Literal)
			return tok
		} else {
			tok = token.NewTokenFromCharacter(token.ILLEGAL, l.ch, l.loc)
		}
	}

	l.readChar()

	return tok
}

// Reads the next char and advances our position in input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.loc.Column += 1
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readEndOfLine() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '\n' || l.ch == 0 {
			break
		}
	}
	l.loc.Line += 1
	l.loc.Column = 0
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	p := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[p:l.position]
}

func (l *Lexer) readNumber() string {
	p := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[p:l.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for {
		switch l.ch {
		case '\r', '\n':
			l.loc.Line += 1
			l.loc.Column = 0
			fallthrough
		case ' ', '\t':
			l.readChar()
		default:
			return
		}
	}
}
