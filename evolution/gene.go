package evolution

import "strconv"

const (
	GeneLength = 4 // GeneLength is the number of bits that represent a gene string.

	EncodedGeneZero     = "0000"
	EncodedGeneOne      = "0001"
	EncodedGeneTwo      = "0010"
	EncodedGeneThree    = "0011"
	EncodedGeneFour     = "0100"
	EncodedGeneFive     = "0101"
	EncodedGeneSix      = "0110"
	EncodedGeneSeven    = "0111"
	EncodedGeneEight    = "1000"
	EncodedGeneNine     = "1001"
	EncodedGenePlus     = "1010"
	EncodedGeneMinus    = "1011"
	EncodedGeneMultiply = "1100"
	EncodedGeneDivide   = "1101"
)

var (
	// geneTable is a lookup table of possible encoded string gene values to decoded equation
	// part.
	geneTable = map[string]string{
		EncodedGeneZero:     "0",
		EncodedGeneOne:      "1",
		EncodedGeneTwo:      "2",
		EncodedGeneThree:    "3",
		EncodedGeneFour:     "4",
		EncodedGeneFive:     "5",
		EncodedGeneSix:      "6",
		EncodedGeneSeven:    "7",
		EncodedGeneEight:    "8",
		EncodedGeneNine:     "9",
		EncodedGenePlus:     "+",
		EncodedGeneMinus:    "-",
		EncodedGeneMultiply: "*",
		EncodedGeneDivide:   "/",
	}
)

type (
	// GeneDecoder is an interface that allows for the decoding of genetic data into a
	// equation part.
	GeneDecoder interface {
		Decode() string
		IsNumeric() bool
		IsOperator() bool
	}

	// StringGene represents a single gene that has an encoded string representation.
	StringGene struct {
		Encoded string
	}

	// ByteGene represents a single gene that has an encoded byte representation.
	ByteGene struct {
		Encoded byte
	}
)

// Decode gets the decoded equation part of the gene's encoded string.
func (g *StringGene) Decode() string {
	if v, ok := geneTable[g.Encoded]; ok {
		return v
	}
	return " "
}

// IsNumeric returns true if the gene is a numeric value.
func (g *StringGene) IsNumeric() bool {
	value, _ := strconv.ParseInt(g.Encoded, 2, 32)
	isNumeric := (value < 10)
	return isNumeric
}

// IsOperator returns true if the gene is an operator.
func (g *StringGene) IsOperator() bool {
	value, _ := strconv.ParseInt(g.Encoded, 2, 32)
	return value > 9 && value < 14
}

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
