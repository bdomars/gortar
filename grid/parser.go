package grid

import (
	"errors"
	"strconv"
)

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

func (p *parser) parseNumber(single bool) (uint8, error) {
	buffer := make([]byte, 0, 2)
	for {
		t, err := p.takeOne()
		if err != nil {
			return 0, err
		}

		if t.kind == NUMBER {
			buffer = append(buffer, t.data)
		}

		if t.kind == SEPARATOR {
			break
		}

		if len(buffer) == 2 {
			break
		}

		if single {
			break
		}
	}
	num, err := strconv.ParseUint(string(buffer), 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(num), nil
}

func (g Grid) Parse(input string) (GridRef, error) {
	p := newParser(input)

	if letter, err := p.parseLetter(); err == nil {
		p.result.Letter = letter
	} else {
		return p.result, err
	}

	if major, err := p.parseNumber(false); err == nil {
		p.result.Major = major
	} else {
		return p.result, err
	}

	for {
		kp, err := p.parseNumber(true)
		if errors.Is(err, EndOfData{}) {
			break
		} else if err != nil {
			return p.result, err
		}
		p.result.Keypads = append(p.result.Keypads, kp)
	}

	p.result.Grid = &g
	return p.result, nil
}
