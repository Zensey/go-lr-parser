package parser

import (
	"fmt"
)

type TokRec struct {
	Tok      Token
	Txt      string
	Children Stack
}

type Stack []*TokRec

func newTokRec(tok Token) TokRec {
	return TokRec{
		Tok:      tok,
		Children: nil,
	}
}

func newTokRecStr(tok Token, txt string) TokRec {
	return TokRec{
		Tok:      tok,
		Txt:      txt,
		Children: nil,
	}
}

func newTokRecChildren(tok Token, children Stack) TokRec {
	return TokRec{
		Tok:      tok,
		Children: children,
	}
}

func (t TokRec) dump() string {
	ch := ""
	if len(t.Children) > 0 {
		ch += ""
		for i, v := range t.Children {
			ch += v.dump()
			if i < len(t.Children)-1 {
				ch += ", "
			}
		}
	}
	return fmt.Sprintf("%s(%s%s)", t.Tok.String(), t.Txt, ch)
}

//////////////////////////////////////////////////////////////////////
func (s *Stack) clear() {
	*s = (*s)[:0]
}

func (s *Stack) Push(t TokRec) {
	*s = append(*s, &t)
}

func (s *Stack) Pop(n int) Stack {
	l := len(*s)

	popped := make(Stack, n)
	copy(popped, (*s)[l-n:])

	*s = (*s)[0 : l-n]

	return popped
}

func (s Stack) cmpLastTok(t []Token) bool {
	if len(t) > len(s) {
		return false
	}
	for i := 0; i < len(t); i++ {
		if s[len(s)-len(t)+i].Tok != t[i] {
			return false
		}
	}
	return true
}

func (s Stack) peek() Token {
	return s[len(s)-1].Tok
}

func (s Stack) dumpStr() (str string) {
	for _, v := range s {
		str += fmt.Sprintf(v.dump())
	}
	return
}
