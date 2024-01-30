package evolution

import (
	"math/rand"
	"fmt"
	"strconv"
	"strings"
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

// Mutate the StringGene by iterating over the string "bits" and randomly inverting them
// according to the given rate.
func (g *StringGene) Mutate(r float64) {
	var sb strings.Builder
	for _, bit := range g.Encoded {
		mutate := rand.Intn(1000) <= int(r*1000)
		if mutate {
			if string(bit) == "0" {
				sb.WriteString("1")
			} else {
				sb.WriteString("0")
			}
		} else {
			sb.WriteRune(bit)
		}
	}
	g.Encoded = sb.String()
}

// Crossover this StringGene with another StringGene at the provided index to create a new 
// StringGene.
func (g *StringGene) Crossover(o GeneDecoder, i int) GeneDecoder {
	var sb strings.Builder
	sb.WriteString(g.Encoded[:i])
	sb.WriteString(o.String()[i:])
	return &StringGene{sb.String()}
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

// Mutate the ByteGene by creating a bitmask and applying it to the encoded byte
// with an XOR operation.
func (g *ByteGene) Mutate(r float64) {
	var sb strings.Builder
	for i:= 0; i < 3; i++ {
		mutate := rand.Intn(1000) <= int(r*1000)
		if mutate {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	mask, _ := strconv.ParseInt(sb.String(), 2, 32)
	g.Encoded = g.Encoded ^ byte(mask)
}

// Crossover this ByteGene with another GeneDecoder at the provided index to create a new 
// ByteGene.
func (g *ByteGene) Crossover(o GeneDecoder, i int) GeneDecoder {
	var ob *ByteGene
	if i > GeneLength - 1 {
		panic("Crossover index out of range")
	}
	switch v := o.(type) {
	case *StringGene:
		b, _ := strconv.ParseInt(v.String(), 2, 4)
		ob = NewByteGene(byte(b))
	case *ByteGene:
		ob = v
	}
	// create new byte gene with bitwise operations
	// 1. create a bitmask for o
	// 2. create the crossover bits using the bitmask
	// 3. calculate how many bits to shift this gene by in the crossover
	// 4. shift this gene and apply the crossover bits 
	mask := byte(0x0F) >> i
	cBits := ob.Encoded & mask
	shift := byte(i * -1 + 4)
	n := (g.Encoded >> shift << shift) | (cBits)
	return NewByteGene(n)
}
