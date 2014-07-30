package regex

import (
	"fmt"
	"strings"
)

type Parser struct {
	lex       *Lex
	handlers  []Handler
	nextToken *Token
}

func NewParser(lex *Lex) *Parser {
	return &Parser{
		lex:       lex,
		nextToken: lex.GetToken(),
	}
}

func (parser *Parser) Parse() ([]Handler, error) {
	defer func() ([]Handler, error) {
		if e := recover(); e != nil {
			return nil, e.(error)
		}
		return parser.handlers, nil
	}()
	parser.expr()
	return parser.handlers, nil
}

func (p *Parser) error(tt TokenType) {
	fmt.Printf("INVALID Regular expression, Type: %s, expected: %s\n", TokenTypeMap[p.nextToken.Type], TokenTypeMap[tt])
	leftspace := strings.Repeat(" ", p.lex.Cur-1)
	rightSpace := strings.Repeat(" ", len(p.lex.Source)-p.lex.Cur)
	s := fmt.Errorf("%s\n%s%c%s", p.lex.Source, leftspace, '^', rightSpace)
	fmt.Println(s.Error())
	panic(s)
}

func (parser *Parser) consume(tt TokenType) {
	if parser.nextToken.Type != tt {
		parser.error(tt)
	}
	parser.nextToken = parser.lex.GetToken()
}

func (parser *Parser) expr() {
	parser.term()
	if parser.nextToken.Type == ALT {
		tok := parser.nextToken
		parser.consume(ALT)
		parser.expr()
		parser.handlers = append(parser.handlers, NewAltHandler(tok))
	}
}

func (parser *Parser) term() {
	parser.factor()
	if parser.nextToken.Type != EOF && parser.nextToken.Type != ALT {
		parser.term()
		parser.handlers = append(parser.handlers, NewConsHandler(NewToken(CONS, ' ')))
	}
}

func (parser *Parser) factor() {
	parser.primary()
	if parser.nextToken.Type == STAR || parser.nextToken.Type == PLUS {
		parser.handlers = append(parser.handlers, NewReplHandler(parser.nextToken))
		parser.consume(parser.nextToken.Type)
	}
}

func (parser *Parser) primary() {
	if parser.nextToken.Type == LP {
		parser.consume(LP)
		parser.expr()
		parser.consume(RP)
		return
	}
	if parser.nextToken.Type == CHAR {
		parser.handlers = append(parser.handlers, NewCharHandler(parser.nextToken))
		parser.consume(CHAR)
		return
	}
	if parser.nextToken.Type == QMARK {
		parser.handlers = append(parser.handlers, NewQMarkHandler(parser.nextToken))
		parser.consume(QMARK)
		return
	}
	if parser.nextToken.Type != EOF {
		parser.error(EOF)
	}
}
