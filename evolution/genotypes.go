package evolution

import (
	"fmt"
	"strings"
)

const (
	ChromosomeLength = 9 // ChromosomeLength is the number of Genes contained in each Chromosome.
)

type (
	// GeneDecoder defines necessary functions for decoding a Gene.
	GeneDecoder interface {
		fmt.Stringer
		Decode() string
		IsNumeric() bool
		IsOperator() bool
	}

	// GeneMutator defines necessary functions for mutating a Gene.
	GeneMutator interface {
		Mutate(r float64)
	}

	// GeneDecoderMutator defines necessary functions for decoding and mutating a Gene.
	GeneDecoderMutator interface {
		GeneDecoder
		GeneMutator
	}
)

// Genotype contains a Chromosome - a list of genes that make up a potential
// solution to a problem.
type Genotype struct {
	Chromosome [ChromosomeLength]GeneDecoderMutator
}

// ToEncodedString converts the Genotype to an encoded string of bits.
func (g *Genotype) ToEncodedString() string {
	var sb strings.Builder
	for _, gene := range g.Chromosome {
		sb.WriteString(gene.String())
	}
	return sb.String()
}

// ToDecodedString converts the Genotype to a decoded string of values.
func (g *Genotype) ToDecodedString() string {
	var sb strings.Builder
	for _, gene := range g.Chromosome {
		sb.WriteString(gene.Decode())
		sb.WriteString(" ")
	}
	return sb.String()
}

// ToFormula converts the Genotype to a proper formula after discarding
// nonsensical data.
func (g *Genotype) ToFormula() string {
	var sb strings.Builder
	haveNumeric := false
	for i, gene := range g.Chromosome {
		if haveNumeric {
			if gene.IsOperator() && i < ChromosomeLength-1 && containsNumeric(g.Chromosome[i:]) {
				sb.WriteString(gene.Decode())
				sb.WriteString(" ")
				haveNumeric = false
			}
		} else {
			if gene.IsNumeric() {
				sb.WriteString(gene.Decode())
				sb.WriteString(" ")
				haveNumeric = true
			}
		}
	}
	return strings.TrimSpace(sb.String())
}

// Mutate the Genotype by iterating over all bits of the EncodedString and
// randomly flipping bits according to the mutationRate.
func (g *Genotype) Mutate(r float64) *Genotype {
	for _, gene := range g.Chromosome {
		gene.Mutate(r)
	}
	return g
}

// Crossover this Genotype with another Genotype to create two new Genotypes.
func (g *Genotype) Crossover(other *Genotype) (*Genotype, *Genotype) {

}

// Determines if the gene slice contains a numeric decoded value
func containsNumeric(genes []GeneDecoderMutator) bool {
	for _, gene := range genes {
		if gene.IsNumeric() {
			return true
		}
	}
	return false
}
