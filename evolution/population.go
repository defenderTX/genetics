package evolution

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Population struct {
	members     []*StringGenotype
	Generations uint32
}

func (p *Population) String() string {
	var sb strings.Builder
	for _, genotype := range p.members {
		sb.WriteString(fmt.Sprintf("%s : %s : %d\n", genotype, genotype.ToFormula(), SolveExpression(genotype.ToFormula())))
	}
	return sb.String()
}

func (p *Population) ContainsSolution(target int) (bool, *StringGenotype) {
	for _, genotype := range p.members {
		if SolveExpression(genotype.ToFormula()) == target {
			return true, genotype
		}
	}
	return false, nil
}

func GenerateInitialPopulation(size int) *Population {
	members := []*StringGenotype{}
	for i := 0; i < size; i++ {
		members = append(members, generateRandomGenotype())
	}
	return &Population{
		members: members, 
		Generations: 1,
	}
}

func generateRandomGenotype() *StringGenotype {
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

func EvolvePopulation(currentPopulation *Population,
	target int,
	crossoverRate float64,
	mutationRate float64) *Population {
	members := []*StringGenotype{}
	for i := 0; i < len(currentPopulation.members)/2; i++ {
		genotype1, genotype2 := selectFittest(target, currentPopulation)
		genotype1, genotype2 = applyCrossover(genotype1, genotype2)
		genotype1 = genotype1.Mutate(float64(0.001))
		genotype2 = genotype2.Mutate(float64(0.001))
		members = append(members, genotype1)
		members = append(members, genotype2)
	}
	nextPopulation := &Population{members, currentPopulation.Generations + 1}
	return nextPopulation
}

func selectFittest(target int, currentPopulation *Population) (*StringGenotype, *StringGenotype) {
	rouletteWheel := []*StringGenotype{}
	for _, genotype := range currentPopulation.members {
		fitness := determineFitness(target, genotype)
		slices := int(fitness * 100)
		for i := 0; i < slices; i++ {
			rouletteWheel = append(rouletteWheel, genotype)
		}
	}
	var genotype1, genotype2 *StringGenotype
	genotype1 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	genotype2 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	return genotype1, genotype2
}

func determineFitness(target int, genotype *StringGenotype) float64 {
	return float64(1.0) / (float64(target) - float64(SolveExpression(genotype.ToFormula())))
}

func applyCrossover(genotype1, genotype2 *StringGenotype) (*StringGenotype, *StringGenotype) {
	g1 := genotype1
	g2 := genotype2
	crossoverRate := float64(0.7)
	if rand.Intn(101) <= int(crossoverRate*100) {
		crossoverAt := rand.Intn(ChromosomeLength * GeneLength)
		g1 = genotype1.Crossover(genotype2, crossoverAt)
		g2 = genotype2.Crossover(genotype1, crossoverAt)
	}
	return g1, g2
}
