package evolution

import (
	"math/rand"
	"strings"
)

const (
	// ChromosomeLength is the number of Genes contained in each Chromosome
	ChromosomeLength = 9
)

// Genotype contains a Chromosome - a list of genes that make up a potential
// solution to a problem
type Genotype struct {
	Chromosome []StringGene
}

// ToEncodedString converts the Genotype to an encoded string of bits
func (genotype Genotype) ToEncodedString() string {
	encodedString := ""
	for _, gene := range genotype.Chromosome {
		encodedString += gene.Encoded
	}
	return encodedString
}

// ToDecodedString converts the Genotype to a decoded string of values
func (genotype Genotype) ToDecodedString() string {
	decodedString := ""
	for _, gene := range genotype.Chromosome {
		decodedString += gene.Decode()
		decodedString += " "
	}
	return decodedString
}

// ToFormula converts the Genotype to a proper formula after discarding
// nonsensical data
func (genotype Genotype) ToFormula() string {
	formulaString := ""
	previousNumeric := false
	for index, gene := range genotype.Chromosome {
		if !previousNumeric && gene.IsNumeric() {
			formulaString += gene.Decode() + " "
			previousNumeric = true
		}
		if previousNumeric && gene.IsOperator() &&
			index < len(genotype.Chromosome)-1 &&
			containsNumeric(genotype.Chromosome[index:]) {
			formulaString += gene.Decode() + " "
			previousNumeric = false
		}
	}
	return strings.TrimSpace(formulaString)
}

// Mutate the Genotype by iterating over all bits of the EncodedString and
// randomly flipping bits according to the mutationRate.
func (genotype Genotype) Mutate(mutationRate float64) Genotype {
	mutatedGenotype := Genotype{[]StringGene{}}
	for _, gene := range genotype.Chromosome {
		var temp string
		for _, bit := range gene.Encoded {
			mutate := rand.Intn(1000) <= int(mutationRate*1000)
			if mutate {
				if string(bit) == "0" {
					temp += "1"
				} else {
					temp += "0"
				}
			} else {
				temp += string(bit)
			}
		}
		mutatedGenotype.Chromosome = append(mutatedGenotype.Chromosome, StringGene{temp})
	}
	return mutatedGenotype
}

// Determines if the gene slice contains a numeric decoded value
func containsNumeric(genes []StringGene) bool {
	for _, gene := range genes {
		if gene.IsNumeric() {
			return true
		}
	}
	return false
}
