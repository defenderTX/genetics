package evolution

import (
	"testing"
)

func TestToFormulaNoNumerics(t *testing.T) {
	noNumerics := Genotype{[]StringGene{plus, minus, multiply, divide, plus, minus}}
	if noNumerics.ToFormula() != "" {
		t.Fail()
	}
}

func TestToFormulaAllNumerics(t *testing.T) {
	allNumerics := Genotype{[]StringGene{zero, one, two, three, four, five, six}}
	if allNumerics.ToFormula() != zero.Decode() {
		t.Fail()
	}
}

func TestMutate(t *testing.T) {
	allFives := Genotype{[]StringGene{}}
	allPluses := Genotype{[]StringGene{}}
	for len(allFives.Chromosome) != ChromosomeLength {
		allFives.Chromosome = append(allFives.Chromosome, five)
		allPluses.Chromosome = append(allPluses.Chromosome, plus)
	}
	if allFives.Mutate(float64(1.000)).ToEncodedString() != allPluses.ToEncodedString() {
		t.Fail()
	}
}
