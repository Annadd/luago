package ast

type Block struct {
	LastLine int
	Stats    []Stat // statement
	RetExps  []Exp  // return expression
}
