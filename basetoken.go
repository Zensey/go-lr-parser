package parser

//go:generate stringer -type=BaseToken

type Token interface {
	String() string
}

type BaseToken int

const (
	TokUndefined BaseToken = iota

	Beg
	End
	None
)
