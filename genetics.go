package main

import (
	"math/rand"
	"os"
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
	target, _ := strconv.Atoi(os.Args[1])
	pop := GenerateRandomPopulation(50)
	for pop.generations < 100 {
		for _, genotype := range pop.members {
			if parser.SolveExpression(genotype.ToFormula()) == target {
				println(genotype.ToFormula())
				os.Exit(0)
			}
		}
		pop = evolvePopulation(target, &pop)
		PrintPopulation(pop)
	}
}

func evolvePopulation(target int, currentPopulation *population) population {
	members := []genotypes.Genotype{}
	for i := 0; i < 25; i++ {
		genotype1, genotype2 := selectFittest(target, currentPopulation)
		genotype1, genotype2 = applyCrossover(genotype1, genotype2)
		members = append(members, genotype1)
		members = append(members, genotype2)
	}
	nextPopulation := population{members, currentPopulation.generations + 1}
	return nextPopulation
}

func selectFittest(target int, currentPopulation *population) (genotypes.Genotype, genotypes.Genotype) {
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
	for genotype1.ToEncodedString() == genotype2.ToEncodedString() {
		genotype2 = rouletteWheel[rand.Intn(len(rouletteWheel))]
	}
	return genotype1, genotype2
}

func determineFitness(target int, genotype genotypes.Genotype) float64 {
	return float64(1.0) / (float64(target) - float64(parser.SolveExpression(genotype.ToFormula())))
}

func applyCrossover(genotype1 genotypes.Genotype, genotype2 genotypes.Genotype) (genotypes.Genotype, genotypes.Genotype) {
	crossoverRate := float64(0.7)
	crossover := rand.Intn(101) >= int(crossoverRate*100)
	if crossover {

	}
}
