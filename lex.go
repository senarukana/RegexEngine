package regex

type Lex struct {
	Cur    int
	Source string
	Runes  []rune
}

func NewLex(s string) *Lex {
	lex := &Lex{
		Source: s,
		Runes:  []rune(s),
	}
	return lex
}

func (lex *Lex) GetToken() *Token {
	if lex.Cur < len(lex.Runes) {
		val := lex.Runes[lex.Cur]
		lex.Cur += 1
		if tt, ok := TokenMap[val]; ok {
			return NewToken(tt, val)
		} else {
			return NewToken(CHAR, val)
		}
	} else {
		return NewToken(EOF, ' ')
	}
}
