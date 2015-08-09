package genotypes

import "github.com/defendertx/genetics/genes"

const (
	ChromosomeLength = 9
)

type Genotype struct {
	Chromosome []genes.Gene
}

func (genotype Genotype) ChromosomeToEncodedString() string {
	encodedString := ""
	for _, gene := range genotype.Chromosome {
		encodedString += gene.EncodedString
	}
	return encodedString
}

func (genotype Genotype) ChromosomeToDecodedString() string {
	decodedString := ""
	for _, gene := range genotype.Chromosome {
		decodedString += gene.DecodedValue()
		decodedString += " "
	}
	return decodedString
}

func (genotype Genotype) ChromosomeToFormula() string {
	formulaString := ""
	previousNumeric := false
	for index, gene := range genotype.Chromosome {
		if !previousNumeric && gene.IsNumeric() {
			formulaString += gene.DecodedValue()
			formulaString += " "
			previousNumeric = true
		}
		if previousNumeric && gene.IsOperator() && index < len(genotype.Chromosome) {
			formulaString += gene.DecodedValue()
			formulaString += " "
			previousNumeric = false
		}
	}
	return formulaString
}
