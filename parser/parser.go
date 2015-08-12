package parser

type tree struct {
	left  *tree
	token string
	right *tree
}
