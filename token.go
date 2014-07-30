package regex

import (
	"fmt"
)

type TokenType int

const (
	INVALID TokenType = iota
	EOF

	CHAR
	PLUS  // +
	QMARK // ?
	STAR  // *
	ALT   // |
	LP    // (
	RP    //)

	CONS
)

var TokenTypeMap = map[TokenType]string{
	INVALID: "INVALID",
	EOF:     "EOF",
	PLUS:    "+",
	STAR:    "*",
	ALT:     "|",
	QMARK:   "?",
	LP:      "LP",
	RP:      "RP",
	CONS:    "CONS",
	CHAR:    "CHAR",
}

var TokenMap = map[rune]TokenType{
	'+': PLUS,
	'*': STAR,
	'|': ALT,
	'?': QMARK,
	'(': LP,
	')': RP,
}

type Token struct {
	Type TokenType
	Val  rune
}

func NewToken(t TokenType, v rune) *Token {
	return &Token{
		Type: t,
		Val:  v,
	}
}

func (tok *Token) String() string {
	return fmt.Sprintf("%c", tok.Val)
}
