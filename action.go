package parser

type Action int

const (
	ActShift Action = iota
	ActReduce
	ActAccept
	ActNone
)
