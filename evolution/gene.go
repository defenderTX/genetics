package evolution

import (
	"fmt"
	"strconv"
)

const (
	GeneLength = 4 // GeneLength is the number of bits that represent a gene string.
)

type (
	// StringGene represents a single gene that has an encoded string representation.
	StringGene struct {
		Encoded string
	}

	// ByteGene represents a single gene that has an encoded byte representation.
	ByteGene struct {
		Encoded byte
	}
)

// NewStringGene initializes and returns a new StringGene from a string.
func NewStringGene(s string) *StringGene {
	return &StringGene{
		Encoded: s,
	}
}

// NewStringGeneFromByte initializes and returns a new StringGene from a byte.
func NewStringGeneFromByte(b byte) *StringGene {
	bits := fmt.Sprintf("%04b", b)
	return &StringGene{
		Encoded: bits[len(bits)-GeneLength:],
	}
}

// Decode gets the decoded equation part of the gene's encoded string.
func (g *StringGene) Decode() string {
	v, _ := strconv.ParseInt(g.Encoded, 2, 32)
	if v < 0xA {
		return strconv.Itoa(int(v))
	}
	switch v {
	case 0xA:
		return "+"
	case 0xB:
		return "-"
	case 0xC:
		return "*"
	case 0xD:
		return "/"
	default:
		return " "
	}
}

// IsNumeric returns true if the gene is a numeric value.
func (g *StringGene) IsNumeric() bool {
	v, _ := strconv.ParseInt(g.Encoded, 2, 32)
	return v < 0xA
}

// IsOperator returns true if the gene is an operator.
func (g *StringGene) IsOperator() bool {
	v, _ := strconv.ParseInt(g.Encoded, 2, 32)
	return v > 0x9 && v < 0xE
}

// String implementation of the Stringer interface for StringGene.
func (g *StringGene) String() string {
	return g.Encoded
}

// NewByteGene initializes and returns a new ByteGene from a byte.
func NewByteGene(b byte) *ByteGene {
	return &ByteGene{
		Encoded: b,
	}
}

// Decoded gets the decoded equation part of the gene's encoded byte.
func (g *ByteGene) Decode() string {
	if g.Encoded < 0xA {
		return strconv.Itoa(int(g.Encoded))
	}
	switch g.Encoded {
	case 0xA:
		return "+"
	case 0xB:
		return "-"
	case 0xC:
		return "*"
	case 0xD:
		return "/"
	default:
		return " "
	}
}

// IsNumeric returns true if the gene is a numeric value.
func (g *ByteGene) IsNumeric() bool {
	return g.Encoded < 0xA
}

// IsOperator returns true if the gene is an operator.
func (g *ByteGene) IsOperator() bool {
	return g.Encoded > 0x9 && g.Encoded < 0xE
}

// String implementation of the Stringer interface for ByteGene.
func (g *ByteGene) String() string {
	return fmt.Sprintf("%04b", g.Encoded)
}
