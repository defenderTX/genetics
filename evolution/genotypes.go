package evolution

import (
	"fmt"
	"math/rand"
	"strconv"
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

// StringGenotype contains a Chromosome made up of StringGenes - a list of genes that make up a potential
// solution to a problem.
type StringGenotype struct {
	Chromosome [ChromosomeLength]*StringGene
}

// NewStringGenotype initializes and returns a new StringGenotype with random data.
func NewStringGenotype() *StringGenotype {
	chromosome := [ChromosomeLength]*StringGene{}
	for i := 0; i < ChromosomeLength; i++ {
		gene := ""
		for j := 0; j < GeneLength; j++ {
			gene += strconv.Itoa(rand.Intn(2))
		}
		chromosome[i] = &StringGene{gene}
	}
	return &StringGenotype{chromosome}
}

// String converts the Genotype to an encoded string of bits.
func (g *StringGenotype) String() string {
	var sb strings.Builder
	for _, gene := range g.Chromosome {
		sb.WriteString(gene.String())
	}
	return sb.String()
}

// ToDecodedString converts the Genotype to a decoded string of values.
func (g *StringGenotype) Decoded() string {
	var sb strings.Builder
	for _, gene := range g.Chromosome {
		sb.WriteString(gene.Decode())
		sb.WriteString(" ")
	}
	return sb.String()
}

// Formula converts the Genotype to a proper formula after discarding
// nonsensical data.
func (g *StringGenotype) Formula() string {
	var sb strings.Builder
	haveNumeric := false
	for i, gene := range g.Chromosome {
		if haveNumeric {
			if gene.IsOperator() && i < ChromosomeLength-1 && g.containsNumeric(g.Chromosome[i:]) {
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
func (g *StringGenotype) Mutate(r float64) *StringGenotype {
	for _, gene := range g.Chromosome {
		gene.Mutate(r)
	}
	return g
}

// Crossover this Genotype with another Genotype at the provided index to create a new Genotype.
func (g *StringGenotype) Crossover(o *StringGenotype, i int) *StringGenotype {
	child := &StringGenotype{}
	geneSkips := int(i / GeneLength)
	for i := 0; i < geneSkips; i++ {
		// add partial chromosome that is unmodified
		child.Chromosome[i] = g.Chromosome[i]
	}
	if i%GeneLength != 0 {
		// handle mid-gene crossover
		bitSkips := i % GeneLength
		g1 := g.Chromosome[geneSkips]
		g2 := o.Chromosome[geneSkips]
		child.Chromosome[geneSkips] = g1.Crossover(g2, bitSkips)
		geneSkips++
	}
	for i := geneSkips; i < ChromosomeLength; i++ {
		// add partial chromosome for crossover
		child.Chromosome[i] = o.Chromosome[i]
	}
	return child
}

// Determines if the gene slice contains a numeric decoded value
func (g *StringGenotype) containsNumeric(genes []*StringGene) bool {
	for _, gene := range genes {
		if gene.IsNumeric() {
			return true
		}
	}
	return false
}
