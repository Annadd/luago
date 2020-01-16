package ast

type Exp interface{}
type NilExp struct{ Line int }
type TrueExp struct{ Line int }
type FalseExp struct{ Line int }
type VarargExp struct{ Line int }

// Numeral
type IntegerExp struct {
	Line int
	Val  int64
}

type FloatExp struct {
	Line int
	Val  float64
}

// LiteralString
type StringExp struct {
	Line int
	Str  string
}

// unop exp
type UnopExp struct {
	Line int // line of operator
	Op   int // operator
	Exp  Exp
}

// exp1 op exp2
type BinopExp struct {
	Line int // line of operator
	Op   int // operator
	Exp1 Exp
	Exp2 Exp
}

type ConcatExp struct {
	Line int // line of last ..
	Exps []Exp
}

type TableConstructorExp struct {
	Line     int // line of `{`
	LastLine int // line of `}`
	KeyExps  []Exp
	ValExps  []Exp
}

// functiondef ::= function funcbody
type FuncDefExp struct {
	Line     int
	LastLine int // line of `end`
	ParList  []string
	IsVararg bool
	Block    *Block
}

type NameExp struct {
	Line int
	Name string
}

type ParensExp struct {
	Exp Exp
}

type TableAccessExp struct {
	LastLine  int // linf of `]`
	PrefixExp Exp
	KeyExp    Exp
}

type FuncCallExp struct {
	Line      int // line of `(`
	LastLine  int // line of `)`
	PrefixExp Exp
	NameExp   *StringExp
	Args      []Exp
}
