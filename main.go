package main

import (
	"fmt"
	"os"

	"github.com/mwillfox/go-eq-gen/evolution"

	"github.com/alecthomas/kong"
)

type (
	CLI struct {
		Target      int `arg:"" help:"Target number to solve for."`
		Generations int `short:"g" default:"10000" help:"Number of generations to run."`
		Size        int `short:"s" default:"100" help:"Size of the population."`
	}
)

func main() {
	cli := &CLI{}
	_ = kong.Parse(cli)
	pop := evolution.NewPopulation(cli.Size)
	for pop.Generations < uint32(cli.Generations) {
		s, ok := pop.Solution(cli.Target)
		if ok {
			fmt.Printf("Solution found in %d generations:\n", pop.Generations)
			fmt.Println(s.Formula())
			os.Exit(0)
		}
		fmt.Printf("%d\t%s\n", pop.Generations, pop.Fittest(cli.Target).Formula())
		pop.Evolve(cli.Target)
	}
	fmt.Printf("Solution not found in %d generations\n", pop.Generations)
}
