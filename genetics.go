package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/defendertx/genetics/genetics"
)

func main() {
	target, _ := strconv.Atoi(os.Args[1])
	pop := genetics.GenerateInitialPopulation(50)
	for pop.Generations < 101 {
		found, solution := pop.ContainsSolution(target)
		if found {
			fmt.Printf("Solution found in %d generations:\n", pop.Generations)
			println(solution.ToFormula())
			os.Exit(0)
		}
		pop = genetics.EvolvePopulation(pop, target, float64(0.7), float64(0.001))
	}
	println("Solution not found in 100 generations")
}
