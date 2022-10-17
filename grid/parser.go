package grid

type kind uint8

const (
	SEPARATOR kind = iota
	LETTER
	NUMBER
)

type token struct {
	data string
	pos  int
	kind
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

func (g Grid) Parse(input string) (GridRef, error) {
	p := newParser(input)

	return p.result, nil
}
