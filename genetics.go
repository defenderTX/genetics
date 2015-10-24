package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/defendertx/genetics/genes"
	"github.com/defendertx/genetics/genotypes"
	"github.com/defendertx/genetics/parser"
)

type population struct {
	members     []genotypes.Genotype
	generations uint32
}

func GenerateRandomPopulation(size int32) population {
	members := make([]genotypes.Genotype, size)
	rand.Seed(int64(time.Now().Unix()))
	for i := range members {
		members[i] = generateRandomGenotype()
	}
	population := population{members, 0}
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

func PrintPopulation(population population) {
	for _, genotype := range population.members {
		print(genotype.ToEncodedString(), " : ")
		print(genotype.ToDecodedString(), " : ")
		print(genotype.ToFormula(), " : ")
		println(parser.SolveExpression(genotype.ToFormula()))
	}
}

func main() {
	pop := GenerateRandomPopulation(50)
	PrintPopulation(pop)
}
