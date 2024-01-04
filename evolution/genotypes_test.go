package evolution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestToFormulaNoNumerics(t *testing.T) {

// 	noNumerics := Genotype{[]StringGene{plus, minus, multiply, divide, plus, minus}}
// 	if noNumerics.ToFormula() != "" {
// 		t.Fail()
// 	}
// }

// func TestToFormulaAllNumerics(t *testing.T) {
// 	allNumerics := Genotype{[]StringGene{zero, one, two, three, four, five, six}}
// 	if allNumerics.ToFormula() != zero.Decode() {
// 		t.Fail()
// 	}
// }

// func TestMutate(t *testing.T) {
// 	allFives := Genotype{[]StringGene{}}
// 	allPluses := Genotype{[]StringGene{}}
// 	for len(allFives.Chromosome) != ChromosomeLength {
// 		allFives.Chromosome = append(allFives.Chromosome, five)
// 		allPluses.Chromosome = append(allPluses.Chromosome, plus)
// 	}
// 	if allFives.Mutate(float64(1.000)).ToEncodedString() != allPluses.ToEncodedString() {
// 		t.Fail()
// 	}
// }

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
