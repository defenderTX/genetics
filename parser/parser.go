package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type tree struct {
	left  *tree
	token string
	right *tree
}

func ParseFormula(formula string) {
	fs := token.NewFileSet()
	tr, _ := parser.ParseExpr(formula)
	ast.Print(fs, tr)
}
