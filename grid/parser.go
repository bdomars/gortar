package grid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	SEPARATOR byte = '-'
)

type parser struct {
	data []byte
	next byte
	len  int
	idx  int
	gr   GridRef
}

func newParser(input string, has_separator bool) parser {
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
	if p.next == SEPARATOR {
		p.advance()
	}
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

func (p *parser) parse_number(big_major bool) (uint8, error) {

	if big_major {
		fmt.Println("parsing big major")
		data := p.data[p.idx : p.idx+2]
		num, err := strconv.ParseUint(string(data), 10, 8)

		if err != nil {
			return 0, SyntaxError{
				pos:   p.idx,
				token: p.next,
			}
		}

		return uint8(num), nil

	} else {
		num, err := strconv.ParseUint(string(p.next), 10, 8)

		if err != nil {
			return 0, SyntaxError{
				pos:   p.idx,
				token: p.next,
			}
		}

		return uint8(num), nil
	}
}

func (g Grid) Parse(input string) (GridRef, error) {

	if len(input) < 1 {
		return GridRef{}, EndOfData{}
	}

	has_separator := false
	if strings.ContainsRune(input, rune(SEPARATOR)) {
		has_separator = true
		fmt.Println("has separator")
	}

	p := newParser(input, has_separator)
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

	num, err := p.parse_number(has_separator)
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

		num, err := p.parse_number(false)
		if err != nil {
			return GridRef{}, err
		}

		p.gr.Keypads = append(p.gr.Keypads, num)
	}

	return p.gr, nil
}
