package parser

import "github.com/defendertx/genetics/genes"

type tree struct {
	left  *tree
	token string
	right *tree
}

func Evaluate(chromosome string) int32 {

}

func buildAst(chromosome string) tree {
	values := ""
	for i := 0; i < len(chromosome); i += genes.GENE_LENGTH {
		values += genes.GeneToValue(chromosome[i : i+GENE_LENGTH])
		values += " "
	}
}

func parseFormula()
