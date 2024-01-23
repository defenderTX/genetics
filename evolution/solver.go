package evolution

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

type (
	// Solver implementation which parses the expression into an Abstract Syntax Tree
	// and solves using recursion.
	ASTSolver struct{}
)

// NewASTSolver initializes and returns a new ASTSolver.
func NewASTSolver() *ASTSolver {
	return &ASTSolver{}
}

// SolveExpression takes a string that represents a mathematical expression
// and returns the solution. Supported operators are +, -, *, and /.
func (s *ASTSolver) Solve(e string) (int, error) {
	expressionTree, err := parser.ParseExpr(e)
	if err != nil {
		return 0, err
	}
	return s.solve(expressionTree), err
}

// Takes the ast format of the expression and solves the expression
func (s *ASTSolver) solve(expression ast.Expr) int {
	switch v := expression.(type) {
	case *ast.BinaryExpr:
		switch v.Op {
		case token.ADD:
			return s.solve(v.X) + s.solve(v.Y)
		case token.SUB:
			return s.solve(v.X) - s.solve(v.Y)
		case token.MUL:
			return s.solve(v.X) * s.solve(v.Y)
		case token.QUO:
			// check to prevent divide by zero panic
			if s.solve(v.Y) == 0 {
				return 0
			}
			return s.solve(v.X) / s.solve(v.Y)
		default:
			return 0
		}
	case *ast.BasicLit:
		value, _ := strconv.Atoi(v.Value)
		return value
	default:
		return 0
	}
}
