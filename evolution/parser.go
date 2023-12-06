package evolution

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// SolveExpression takes a string that represents a mathematical expression
// and returns the solution. Supported operators are +, -, *, and /.
func SolveExpression(expression string) int {
	expressionTree, _ := parser.ParseExpr(expression)
	return solve(expressionTree)
}

// Takes the ast format of the expression and solves the expression
func solve(expression ast.Expr) int {
	switch typedExpr := expression.(type) {
	case *ast.BinaryExpr:
		switch typedExpr.Op {
		case token.ADD:
			return solve(typedExpr.X) + solve(typedExpr.Y)
		case token.SUB:
			return solve(typedExpr.X) - solve(typedExpr.Y)
		case token.MUL:
			return solve(typedExpr.X) * solve(typedExpr.Y)
		case token.QUO:
			// check to prevent divide by zero panic
			if solve(typedExpr.Y) == 0 {
				return 0
			}
			return solve(typedExpr.X) / solve(typedExpr.Y)
		default:
			return 0
		}
	case *ast.BasicLit:
		value, _ := strconv.Atoi(typedExpr.Value)
		return value
	default:
		return 0
	}
}
