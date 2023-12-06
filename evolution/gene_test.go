package evolution

import (
	"testing"
)

var (
	zero      = StringGene{"0000"}
	one       = StringGene{"0001"}
	two       = StringGene{"0010"}
	three     = StringGene{"0011"}
	four      = StringGene{"0100"}
	five      = StringGene{"0101"}
	six       = StringGene{"0110"}
	seven     = StringGene{"0111"}
	eight     = StringGene{"1000"}
	nine      = StringGene{"1001"}
	plus      = StringGene{"1010"}
	minus     = StringGene{"1011"}
	multiply  = StringGene{"1100"}
	divide    = StringGene{"1101"}
	numerics  = []StringGene{zero, one, two, three, four, five, six, seven, eight, nine}
	operators = []StringGene{plus, minus, multiply, divide}
)

func TestIsNumericForNumerics(t *testing.T) {
	for _, gene := range numerics {
		if !gene.IsNumeric() {
			t.Fail()
		}
	}
}

func TestIsNumericForOperators(t *testing.T) {
	for _, gene := range operators {
		if gene.IsNumeric() {
			t.Fail()
		}
	}
}

func TestIsOperatorForNumerics(t *testing.T) {
	for _, gene := range numerics {
		if gene.IsOperator() {
			t.Fail()
		}
	}
}

func TestIsOperatorForOperators(t *testing.T) {
	for _, gene := range operators {
		if !gene.IsOperator() {
			t.Fail()
		}
	}
}
