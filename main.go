package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/mwillfox/go-eq-gen/evolution"

	"github.com/alecthomas/kong"
	"github.com/fatih/color"
)

type (
	// CLI defines the command line arguments for the main entry point.
	CLI struct {
		Target      int  `arg:"" help:"Target number to generate an equation which solves to."`
		Generations int  `short:"g" default:"10000" help:"Number of generations to run."`
		Size        int  `short:"s" default:"100" help:"Size of the population."`
		Verbose     bool `short:"v" default:"false" help:"Verbose output."`
	}
)

func printGen(pop *evolution.Population, target, generations int) {
	var sb strings.Builder
	places := len(fmt.Sprintf("%d", generations))
	genString := fmt.Sprintf("%d", pop.Generations)
	pad := places - len(genString)
	for i := 0; i < pad; i++ {
		sb.WriteString("0")
	}
	sb.WriteString(genString)
	cyan := color.New(color.FgCyan)
	yellow := color.New(color.FgYellow)
	cyan.Printf("[Gen %s]\t", sb.String())
	yellow.Print("fittest=")
	fmt.Println(pop.Fittest(target).Formula())
}

// main entry point for the application.
func main() {
	cli := &CLI{}
	_ = kong.Parse(cli)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	gc := func() evolution.GeneDecoder {
		gene := ""
		for j := 0; j < evolution.GeneLength; j++ {
			gene += strconv.Itoa(rand.Intn(2))
		}
		return evolution.NewStringGene(gene)
	}
	pop := evolution.NewPopulation(cli.Size, evolution.NewASTSolver(), gc)
	for pop.Generations <= uint32(cli.Generations) {
		s, ok := pop.Solution(cli.Target)
		if ok {
			if cli.Verbose {
				printGen(pop, cli.Target, cli.Generations)
			}
			fmt.Printf("Solution found in %d generations:\n", pop.Generations)
			green.Println(s.Formula())
			os.Exit(0)
		}
		if cli.Verbose {
			printGen(pop, cli.Target, cli.Generations)
		}
		pop.Evolve(cli.Target)
	}
	red.Printf("Solution not found in %d generations\n", pop.Generations-1)
}
