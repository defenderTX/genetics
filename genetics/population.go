package genetics

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/defendertx/genetics/genes"
	"github.com/defendertx/genetics/genotypes"
	"github.com/defendertx/genetics/parser"
)

type Population struct {
	members     []genotypes.Genotype
	Generations uint32
}

func (p Population) Print() {
	for _, genotype := range p.members {
		print(genotype.ToEncodedString(), " : ")
		print(genotype.ToDecodedString(), " : ")
		print(genotype.ToFormula(), " : ")
		println(parser.SolveExpression(genotype.ToFormula()))
	}
}

func (p Population) ContainsSolution(target int) (bool, genotypes.Genotype) {
	for _, genotype := range p.members {
		if parser.SolveExpression(genotype.ToFormula()) == target {
			return true, genotype
		}
	}
	return false, genotypes.Genotype{}
}

func GenerateInitialPopulation(size uint32) Population {
	members := make([]genotypes.Genotype, size)
	rand.Seed(int64(time.Now().Unix()))
	for i := range members {
		members[i] = generateRandomGenotype()
	}
	population := Population{members, 1}
	return population
}

func generateRandomGenotype() genotypes.Genotype {
	chromosome := []genes.Gene{}
	for i := 0; i < genotypes.ChromosomeLength; i++ {
		gene := ""
		for j := 0; j < genes.GeneLength; j++ {
			gene += strconv.Itoa(rand.Intn(2))
		}
		chromosome = append(chromosome, genes.Gene{gene})
	}
	return genotypes.Genotype{chromosome}
}

func EvolvePopulation(currentPopulation Population,
	target int,
	crossoverRate float64,
	mutationRate float64) Population {
	members := []genotypes.Genotype{}
	for i := 0; i < len(currentPopulation.members)/2; i++ {
		genotype1, genotype2 := selectFittest(target, currentPopulation)
		genotype1, genotype2 = applyCrossover(genotype1, genotype2)
		genotype1 = mutate(genotype1)
		genotype2 = mutate(genotype2)
		members = append(members, genotype1)
		members = append(members, genotype2)
	}
	nextPopulation := Population{members, currentPopulation.Generations + 1}
	return nextPopulation
}

func selectFittest(target int, currentPopulation Population) (genotypes.Genotype, genotypes.Genotype) {
	rouletteWheel := []genotypes.Genotype{}
	for _, genotype := range currentPopulation.members {
		fitness := determineFitness(target, genotype)
		slices := int(fitness * 100)
		for i := 0; i < slices; i++ {
			rouletteWheel = append(rouletteWheel, genotype)
		}
	}
	var genotype1, genotype2 genotypes.Genotype
	genotype1 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	genotype2 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	return genotype1, genotype2
}

func determineFitness(target int, genotype genotypes.Genotype) float64 {
	return float64(1.0) / (float64(target) - float64(parser.SolveExpression(genotype.ToFormula())))
}

func applyCrossover(genotype1 genotypes.Genotype, genotype2 genotypes.Genotype) (genotypes.Genotype, genotypes.Genotype) {
	crossoverRate := float64(0.7)
	crossover := rand.Intn(101) <= int(crossoverRate*100)
	if crossover {
		crossoverAt := rand.Intn(genotypes.ChromosomeLength * genes.GeneLength)
		geneSkips := int(crossoverAt / genes.GeneLength)
		if crossoverAt%genes.GeneLength != 0 {
			bitSkips := crossoverAt % genes.GeneLength
			firstCrossoverGene1 := genotype1.Chromosome[geneSkips].EncodedString
			firstCrossoverGene2 := genotype2.Chromosome[geneSkips].EncodedString
			var temp1, temp2 string
			for i := 0; i < bitSkips; i++ {
				temp1 += string(firstCrossoverGene1[i])
				temp2 += string(firstCrossoverGene2[i])
			}
			for i := bitSkips; i < genes.GeneLength; i++ {
				temp1 += string(firstCrossoverGene2[i])
				temp2 += string(firstCrossoverGene1[i])
			}
			genotype1.Chromosome[geneSkips] = genes.Gene{temp1}
			genotype2.Chromosome[geneSkips] = genes.Gene{temp2}
			geneSkips++
		}
		for i := geneSkips; i < genotypes.ChromosomeLength; i++ {
			temp := genotype1.Chromosome[i]
			genotype1.Chromosome[i] = genotype2.Chromosome[i]
			genotype2.Chromosome[i] = temp
		}
	}
	return genotype1, genotype2
}

func mutate(genotype genotypes.Genotype) genotypes.Genotype {
	mutationRate := float64(0.001)
	for _, gene := range genotype.Chromosome {
		var temp string
		for _, bit := range gene.EncodedString {
			mutate := rand.Intn(1001) <= int(mutationRate*1000)
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
		gene.EncodedString = temp
	}
	return genotype
}
