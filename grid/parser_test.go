package grid

import (
	"regexp"
	"testing"
)

type parseTest struct {
	input    string
	expected GridRef
}

var parseTests = []parseTest{
	{
		"A1",
		GridRef{
			Letter: 'A',
			Major:  1,
		},
	},
	{
		"A12",
		GridRef{
			Letter: 'A',
			Major:  12,
		},
	},
	{
		"A123",
		GridRef{
			Letter: 'A',
			Major:  12,
			Keypads: []uint8{
				3,
			},
		},
	},
	{
		"A1-23",
		GridRef{
			Letter: 'A',
			Major:  1,
			Keypads: []uint8{
				2,
				3,
			},
		},
	},
	{
		"A12-3",
		GridRef{
			Letter: 'A',
			Major:  12,
			Keypads: []uint8{
				3,
			},
		},
	},
}

func TestParser(t *testing.T) {
	g := NewGrid(300)
	for _, test := range parseTests {
		if output, _ := g.Parse(test.input); !output.EqualTo(test.expected) {
			t.Errorf("The grid %s is parsed wrong, expected = %s, got = %s\n", test.input, test.expected, output)
		}
	}
}

func BenchmarkParser(b *testing.B) {
	test_str := "A12-3456789"
	g := NewGrid(300)
	for i := 0; i < b.N; i++ {
		g.Parse(test_str)
	}
}

func BenchmarkRegexpParser(b *testing.B) {
	test_str := "A12-3456789"

	p, _ := regexp.Compile(`^([A-Z])(\d{1,2})[-.,]?(\d*)`)
	for i := 0; i < b.N; i++ {
		p.Find([]byte(test_str))
	}
}
