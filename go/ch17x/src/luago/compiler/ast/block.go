package ast

// chunk ::= block
// type Chunk *Block

// block ::= {stat} [retstat]
type Block struct {
	LastLine int
	Stats    []Stat // statement
	RetExps  []Exp  // return expression
}
