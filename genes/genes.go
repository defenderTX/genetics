package genes

var GENE_LENGTH = 4

// Converts a gene to an ascii value
func GeneToValue(gene string) string {
	switch gene {
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
