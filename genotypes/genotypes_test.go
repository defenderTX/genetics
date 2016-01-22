package genotypes

import (
	"testing"

	"github.com/defendertx/genetics/genes"
)

var (
	zero     = genes.Gene{"0000"}
	one      = genes.Gene{"0001"}
	two      = genes.Gene{"0010"}
	three    = genes.Gene{"0011"}
	four     = genes.Gene{"0100"}
	five     = genes.Gene{"0101"}
	six      = genes.Gene{"0110"}
	seven    = genes.Gene{"0111"}
	eight    = genes.Gene{"1000"}
	nine     = genes.Gene{"1001"}
	plus     = genes.Gene{"1010"}
	minus    = genes.Gene{"1011"}
	multiply = genes.Gene{"1100"}
	divide   = genes.Gene{"1101"}
)

func TestToFormulaNoNumerics(t *testing.T) {
	noNumerics := Genotype{[]genes.Gene{plus, minus, multiply, divide, plus, minus}}
	if noNumerics.ToFormula() != "" {
		t.Fail()
	}
}

func TestToFormulaAllNumerics(t *testing.T) {
	allNumerics := Genotype{[]genes.Gene{zero, one, two, three, four, five, six}}
	if allNumerics.ToFormula() != zero.ToDecodedValue() {
		t.Fail()
	}
}

func TestMutate(t *testing.T) {
	allFives := Genotype{[]genes.Gene{}}
	allPluses := Genotype{[]genes.Gene{}}
	for len(allFives.Chromosome) != ChromosomeLength {
		allFives.Chromosome = append(allFives.Chromosome, five)
		allPluses.Chromosome = append(allPluses.Chromosome, plus)
	}
	if allFives.Mutate(float64(1.000)).ToEncodedString() != allPluses.ToEncodedString() {
		t.Fail()
	}
}
