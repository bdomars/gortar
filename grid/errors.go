package grid

import "fmt"

type EndOfData struct{}

func (e EndOfData) Error() string {
	return "out of data: the parser has run out of data to parse"
}

type SyntaxError struct {
	pos   int
	token byte
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error: failed to parse token '%c' at position %d", e.token, e.pos+1)
}
