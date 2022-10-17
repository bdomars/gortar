package grid

type kind uint8

const (
	SEPARATOR kind = iota
	LETTER
	NUMBER
)

type token struct {
	data byte
	pos  int
	kind kind
}

type parser struct {
	data   []byte
	pos    int
	result GridRef
}

func newParser(input string) parser {
	data := []byte(input)
	return parser{
		data:   data,
		pos:    0,
		result: GridRef{},
	}
}

func (p *parser) takeOne() (token, error) {

	if p.pos == len(p.data) {
		return token{}, EndOfData{}
	}

	var kind kind
	data := p.data[p.pos]
	switch {
	case data > 'A' && data < 'Z':
		kind = LETTER
	case data >= 0 && data <= 9:
		kind = NUMBER
	case data == '-' || data == '.' || data == ',':
		kind = SEPARATOR
	}

	p.pos++

	return token{
		data: data,
		kind: kind,
		pos:  p.pos,
	}, nil

}

func (p *parser) parseLetter() (byte, error) {
	t, err := p.takeOne()
	if err != nil {
		return 0, err
	}

	if t.kind == LETTER {
		return t.data, nil
	}

	return 0, SyntaxError{
		pos:   t.pos,
		token: t.data,
	}
}

func (g Grid) Parse(input string) (GridRef, error) {
	p := newParser(input)

	p.result.Grid = &g
	return p.result, nil
}
