package genes

import "strconv"

const (
	// GeneLength is the number of bits that represent a gene string
	GeneLength = 4
)

// Gene represents a single gene that has an encoded string representation
// that corresponds to a decoded value
type Gene struct {
	EncodedString string
}

// DecodedValue gets the decoded value of the gene's encoded string
func (gene Gene) DecodedValue() string {
	switch gene.EncodedString {
	case "0000":
		return ("0")
	case "0001":
		return ("1")
	case "0010":
		return ("2")
	case "0011":
		return ("3")
	case "0100":
		return ("4")
	case "0101":
		return ("5")
	case "0110":
		return ("6")
	case "0111":
		return ("7")
	case "1000":
		return ("8")
	case "1001":
		return ("9")
	case "1010":
		return ("+")
	case "1011":
		return ("-")
	case "1100":
		return ("*")
	case "1101":
		return ("/")
	default:
		return (" ")
	}
}

func (gene Gene) IsNumeric() bool {
	value, _ := strconv.ParseInt(gene.EncodedString, 2, 4)
	isNumeric := value < 10
	return isNumeric
}

func (gene Gene) IsOperator() bool {
	value, _ := strconv.ParseInt(gene.EncodedString, 2, 4)
	return value > 9 && value < 14
}
