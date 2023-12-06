package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mwillfox/go-eq-gen/evolution"
)

func main() {
	target, _ := strconv.Atoi(os.Args[1])
	pop := evolution.GenerateInitialPopulation(50)
	for pop.Generations <= 10000 {
		found, solution := pop.ContainsSolution(target)
		if found {
			fmt.Printf("Solution found in %d generations:\n", pop.Generations)
			println(solution.ToFormula())
			os.Exit(0)
		}
		pop = evolution.EvolvePopulation(pop, target, float64(0.7), float64(0.001))
	}
	println("Solution not found in 10000 generations")
}
