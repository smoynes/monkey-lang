package ast

import (
	"monkey/token"
)

// Base interface for every node in AST
type Node interface {
	TokenLiteral() string
}

// Statement nodes
type Statement interface {
	Node
	statementNode()
}

// Expression nodes evaluate to a value.
type Expression interface {
	Node
	expressionNode()
}

// Program node is the root of the AST.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Let statements bind a value to a name.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (let *LetStatement) statementNode()       {}
func (let *LetStatement) TokenLiteral() string { return let.Token.Literal }

// Identifiers expressions reference a value.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Return statements return an expression from a function.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
