package ast

type Stat interface{}
type EmptyStat struct{}              // `;`
type BreakStat struct{ Line int }    // break
type LabelStat struct{ Name string } // `::` Name `::`
type GotoStat struct{ Name string }  // goto Name
type DoStat struct{ Block *Block }   // do block end
type FuncCallStat = FuncCallExp      // function call

// if exp then block {elseif exp then block} [else block] end
type IfStat struct {
	Exps   []Exp
	Blocks []*Block
}

// while exp do block end
type WhileStat struct {
	Exp   Exp
	Block *Block
}

// repeat block until exp
type RepeatStat struct {
	Block *Block
	Exp   Exp
}

type ForNumStat struct {
	LineOfFor int
	LineOfDo  int
	VarName   string
	InitExp   Exp
	LimitExp  Exp
	StepExp   Exp
	Block     *Block
}

// for namelist in explist do block end
type ForInStat struct {
	LineOfDo int
	NameList []string
	ExpList  []Exp
	Block    *Block
}

type AssignStat struct {
	LastLine int
	VarList  []Exp
	ExpList  []Exp
}

type LocalVarDeclStat struct {
	LastLine int
	NameList []string
	ExpList  []Exp
}

// local function Name funcbody
type LocalFuncDefStat struct {
	Name string
	Exp  *FuncDefExp
}
