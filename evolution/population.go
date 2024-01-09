package evolution

import (
	"fmt"
	"math/rand"
	"strings"
)

// Population contains the members of a population and the age of the population
// in generations.
type Population struct {
	members     []*StringGenotype                  // members of the population
	fitFn       func(int, *StringGenotype) float64 // fitFn is the fitness function for the population
	cR          float64                            // cR is the crossover rate for the population
	mR          float64                            // mR is the mutation rate for the population
	Generations uint32                             // Generations in the current population
}

// NewPopulation initializes and returns a new population with the given size
// from random data.
func NewPopulation(s int) *Population {
	members := []*StringGenotype{}
	for i := 0; i < s; i++ {
		members = append(members, NewStringGenotype())
	}
	return &Population{
		members: members,
		fitFn: func(t int, g *StringGenotype) float64 {
			return float64(1.0) / (float64(t) - float64(SolveExpression(g.Formula())))
		},
		cR:          float64(0.7),
		mR:          float64(0.001),
		Generations: 1,
	}
}

// String is the fmt.Stringer implementation for Population.
// Returns the population members as a formatted string.
func (p *Population) String() string {
	var sb strings.Builder
	for _, genotype := range p.members {
		sb.WriteString(fmt.Sprintf("%s : %s : %d\n", genotype, genotype.Formula(), SolveExpression(genotype.Formula())))
	}
	return sb.String()
}

// Solution returns the first solution found in the population for the given target
// and a boolean indicating whether a solution was found.
func (p *Population) Solution(t int) (*StringGenotype, bool) {
	for _, g := range p.members {
		if SolveExpression(g.Formula()) == t {
			return g, true
		}
	}
	return nil, false
}

// Fittest returns the fittest member of the population for the given target.
func (p *Population) Fittest(t int) *StringGenotype {
	fittest := p.members[0]
	for _, g := range p.members {
		if p.fitFn(t, g) > p.fitFn(t, fittest) {
			fittest = g
		}
	}
	return fittest
}

// Evolve evolves the population by applying selection, crossover, and mutation using the given
// target, crossover rate, and mutation rate.
func (p *Population) Evolve(t int) {
	members := []*StringGenotype{}
	for i := 0; i < len(p.members)/2; i++ {
		genotype1, genotype2 := p.selectFittest(t)
		genotype1, genotype2 = p.crossover(genotype1, genotype2)
		genotype1 = genotype1.Mutate(p.mR)
		genotype2 = genotype2.Mutate(p.mR)
		members = append(members, genotype1)
		members = append(members, genotype2)
	}
	p.members = members
	p.Generations++
}

// selectFittest returns the two fittest members of the population using roulette wheel selection.
func (p *Population) selectFittest(t int) (*StringGenotype, *StringGenotype) {
	rouletteWheel := []*StringGenotype{}
	for _, genotype := range p.members {
		fitness := p.fitFn(t, genotype)
		slices := int(fitness * 100)
		for i := 0; i < slices; i++ {
			rouletteWheel = append(rouletteWheel, genotype)
		}
	}
	if len(rouletteWheel) > 2 {
		// rouletteWheel must have at least 2 members
		x := rand.Intn(len(rouletteWheel))
		y := rand.Intn(len(rouletteWheel))
		return rouletteWheel[x], rouletteWheel[y]
	} else {
		x := rand.Intn(len(p.members))
		y := rand.Intn(len(p.members))
		return p.members[x], p.members[y]
	}
}

// crossover applies crossover to the given genotypes using the given crossover rate.
func (p *Population) crossover(genotype1, genotype2 *StringGenotype) (*StringGenotype, *StringGenotype) {
	g1 := genotype1
	g2 := genotype2
	if rand.Intn(101) <= int(p.cR*100) {
		crossoverAt := rand.Intn(ChromosomeLength * GeneLength)
		g1 = genotype1.Crossover(genotype2, crossoverAt)
		g2 = genotype2.Crossover(genotype1, crossoverAt)
	}
	return g1, g2
}
