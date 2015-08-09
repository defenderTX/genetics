package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/defendertx/genetics/genes"
)

var CHROMOSOME_LENGTH = 9

type genotype struct {
	chromosome string
}

type population struct {
	members     []genotype
	generations uint32
}

func GenerateRandomPopulation(size int32) population {
	members := make([]genotype, size)
	rand.Seed(int64(time.Now().Unix()))
	for i := range members {
		members[i] = generateRandomGenotype()
	}
	population := population{members, 0}
	return population
}

func generateRandomGenotype() genotype {
	chromosome := ""
	for i := 0; i < genes.GENE_LENGTH*CHROMOSOME_LENGTH; i++ {
		chromosome += strconv.Itoa(rand.Intn(2))
	}
	genotype := genotype{chromosome}
	return genotype
}

func PrintPopulation(population population) {
	for i := range population.members {
		print(population.members[i].chromosome, " : ")
		printGenotype(population.members[i])
	}
}

func printGenotype(genotype genotype) {
	for i := 0; i < CHROMOSOME_LENGTH*genes.GENE_LENGTH; i += genes.GENE_LENGTH {
		printGene(genotype.chromosome[i : i+genes.GENE_LENGTH])
		print(" ")
	}
	println()
}

func printGene(gene string) {
	print(genes.GeneToValue(gene))
}

func main() {
	pop := GenerateRandomPopulation(50)
	PrintPopulation(pop)
}
