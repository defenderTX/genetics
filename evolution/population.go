package evolution

import (
	"math/rand"
	"strconv"
	"time"
)

type Population struct {
	members     []*Genotype
	Generations uint32
}

func (p *Population) Print() {
	for _, genotype := range p.members {
		print(genotype.ToEncodedString(), " : ")
		print(genotype.ToDecodedString(), " : ")
		print(genotype.ToFormula(), " : ")
		println(SolveExpression(genotype.ToFormula()))
	}
}

func (p *Population) ContainsSolution(target int) (bool, *Genotype) {
	for _, genotype := range p.members {
		if SolveExpression(genotype.ToFormula()) == target {
			return true, genotype
		}
	}
	return false, nil
}

func GenerateInitialPopulation(size uint32) *Population {
	members := []*Genotype{}
	for i := range members {
		members[i] = generateRandomGenotype()
	}
	return &Population{
		members: members, 
		Generations: 1,
	}
}

func generateRandomGenotype() *Genotype {
	chromosome := [ChromosomeLength]GeneDecoderMutator{}
	for i := 0; i < ChromosomeLength; i++ {
		gene := ""
		for j := 0; j < GeneLength; j++ {
			gene += strconv.Itoa(rand.Intn(2))
		}
		chromosome[i] = &StringGene{gene}
	}
	return &Genotype{chromosome}
}

func EvolvePopulation(currentPopulation *Population,
	target int,
	crossoverRate float64,
	mutationRate float64) Population {
	members := []*Genotype{}
	for i := 0; i < len(currentPopulation.members)/2; i++ {
		genotype1, genotype2 := selectFittest(target, currentPopulation)
		genotype1, genotype2 = applyCrossover(genotype1, genotype2)
		genotype1 = genotype1.Mutate(float64(0.001))
		genotype2 = genotype2.Mutate(float64(0.001))
		members = append(members, genotype1)
		members = append(members, genotype2)
	}
	nextPopulation := Population{members, currentPopulation.Generations + 1}
	return nextPopulation
}

func selectFittest(target int, currentPopulation *Population) (*Genotype, *Genotype) {
	rouletteWheel := []*Genotype{}
	for _, genotype := range currentPopulation.members {
		fitness := determineFitness(target, genotype)
		slices := int(fitness * 100)
		for i := 0; i < slices; i++ {
			rouletteWheel = append(rouletteWheel, genotype)
		}
	}
	var genotype1, genotype2 *Genotype
	genotype1 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	genotype2 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	return genotype1, genotype2
}

func determineFitness(target int, genotype *Genotype) float64 {
	return float64(1.0) / (float64(target) - float64(SolveExpression(genotype.ToFormula())))
}

func applyCrossover(genotype1, genotype2 *Genotype) (*Genotype, *Genotype) {
	crossoverRate := float64(0.7)
	crossover := rand.Intn(101) <= int(crossoverRate*100)
	if crossover {
		crossoverAt := rand.Intn(ChromosomeLength * GeneLength)
		geneSkips := int(crossoverAt / GeneLength)
		if crossoverAt%GeneLength != 0 {
			bitSkips := crossoverAt % GeneLength
			firstCrossoverGene1 := genotype1.Chromosome[geneSkips].Encoded
			firstCrossoverGene2 := genotype2.Chromosome[geneSkips].Encoded
			var temp1, temp2 string
			for i := 0; i < bitSkips; i++ {
				temp1 += string(firstCrossoverGene1[i])
				temp2 += string(firstCrossoverGene2[i])
			}
			for i := bitSkips; i < GeneLength; i++ {
				temp1 += string(firstCrossoverGene2[i])
				temp2 += string(firstCrossoverGene1[i])
			}
			genotype1.Chromosome[geneSkips] = &StringGene{temp1}
			genotype2.Chromosome[geneSkips] = &StringGene{temp2}
			geneSkips++
		}
		for i := geneSkips; i < ChromosomeLength; i++ {
			temp := genotype1.Chromosome[i]
			genotype1.Chromosome[i] = genotype2.Chromosome[i]
			genotype2.Chromosome[i] = temp
		}
	}
	return genotype1, genotype2
}
