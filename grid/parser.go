package grid

import (
	"errors"
	"strconv"
)

type parser struct {
	data []byte
	next byte
	len  int
	idx  int
	gr   GridRef
}

func newParser(input string) parser {
	data := []byte(input)
	return parser{
		data: data,
		next: data[0],
		idx:  0,
		len:  len(data),
		gr:   GridRef{},
	}
}

func (p *parser) advance() error {
	p.idx = p.idx + 1
	if p.idx == p.len {
		return EndOfData{}
	}
	p.next = p.data[p.idx]
	return nil
}

func (p *parser) parse_letter() (byte, error) {
	if p.next < 'A' || p.next > 'Z' {
		err := SyntaxError{
			token: p.next,
			pos:   p.idx,
		}
		return 0, err
	}
	return p.next, nil
}

func (p *parser) parse_number() (uint8, error) {
	num, err := strconv.ParseUint(string(p.next), 10, 8)
	if err != nil {
		return 0, SyntaxError{
			pos:   p.idx,
			token: p.next,
		}
	}

	return uint8(num), nil
}

func (g Grid) Parse(input string) (GridRef, error) {

	if len(input) < 1 {
		return GridRef{}, EndOfData{}
	}

	p := newParser(input)
	p.gr.Grid = &g

	if letter, err := p.parse_letter(); err != nil {
		return GridRef{}, err
	} else {
		p.gr.Letter = letter
	}

	err := p.advance()
	if err != nil {
		return GridRef{}, errors.New("syntax error: no major number")
	}

	num, err := p.parse_number()
	if err != nil {
		return GridRef{}, err
	}

	p.gr.Major = num

	for {
		err := p.advance()
		if errors.Is(err, EndOfData{}) {
			break
		} else if err != nil {
			return GridRef{}, err
		}

		num, err := p.parse_number()
		if err != nil {
			return GridRef{}, err
		}

		p.gr.Keypads = append(p.gr.Keypads, num)
	}

	return p.gr, nil
}
