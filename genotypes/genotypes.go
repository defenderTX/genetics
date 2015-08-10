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
	validGenes := []genes.Gene{}
	previousNumeric := false
	for index, gene := range genotype.Chromosome {
		if !previousNumeric && gene.IsNumeric() {
			validGenes = append(validGenes, gene)
			previousNumeric = true
		}
		if previousNumeric && gene.IsOperator() && index < len(genotype.Chromosome) {
			validGenes = append(validGenes, gene)
			previousNumeric = false
		}
	}
	return formulaString
}

// Determines if the gene slice contains a numeric decoded value
func containsNumeric(genes []genes.Gene) bool {
	return false
}
