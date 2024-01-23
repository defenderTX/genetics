package evolution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringGenotypeCrossover(t *testing.T) {
	cases := []struct {
		g1 *StringGenotype
		g2 *StringGenotype
		i  int
		e  string
	}{
		{
			g1: &StringGenotype{[ChromosomeLength]*StringGene{
				{"0000"},
				{"0000"},
				{"0000"},
				{"0000"},
				{"0000"},
				{"0000"},
				{"0000"},
				{"0000"},
				{"0000"},
			}},
			g2: &StringGenotype{[ChromosomeLength]*StringGene{
				{"1111"},
				{"1111"},
				{"1111"},
				{"1111"},
				{"1111"},
				{"1111"},
				{"1111"},
				{"1111"},
				{"1111"},
			}},
			i: 14,
			e: "000000000000001111111111111111111111",
		},
	}
	for _, c := range cases {
		g := c.g1.Crossover(c.g2, c.i)
		assert.Equal(t, c.e, g.String(), fmt.Sprintf("%s Crossover(%s, %d) -> %s", c.g1, c.g2, c.i, c.e))
	}
}
