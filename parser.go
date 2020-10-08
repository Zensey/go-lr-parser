package parser

import "fmt"

type ParserTable interface {
	DetectTok(str string) (Token, string, int)
	ActionTable(sym, st Token) (action Action, state Token, nReduce int)
	TableGoto(uncovered, lhs Token) Token
	TableReduceCustom(eState Token, tail Stack) Stack
	TableRecover(s, term Token, st *Stack, termStr string) (Token, string)
}

// https://slideplayer.com/slide/8029737/
// see p24

type Parser struct {
	c           int
	st          Stack
	in          string
	parserTable ParserTable
}

func NewParser(parserTable ParserTable) *Parser {
	return &Parser{
		c:           0,
		st:          make(Stack, 0),
		parserTable: parserTable,
	}
}

func (p *Parser) Reset() {
	p.st.clear()

	p.c = 0
	p.in = ""
	return
}

func (p *Parser) PrintDump() {
	fmt.Println("Dump >", p.st.dumpStr())

	return
}

func (p *Parser) GetState() Stack {
	return p.st
}

func (p *Parser) peekToken() (Token, string) {
	t, res, i := p.parserTable.DetectTok(p.in[p.c:])
	p.c += i

	// fmt.Printf("peekToken > '%v' -> %v i=%d\n", res, t.String(), i)
	return t, res
}

/* Parse parses a line of text */
func (p *Parser) Parse(in string) Action {
	p.in = in
	p.st.Push(newTokRec(Beg))
	term, termStr := p.peekToken()

	for {
		s := p.st.peek()

		eAct, eState, eNPop := p.parserTable.ActionTable(term, s)

		switch eAct {
		case ActShift:
			p.st.Push(newTokRecStr(term, termStr))
			p.st.Push(newTokRec(eState))

			term, termStr = p.peekToken()

		case ActReduce:
			tail := p.st.Pop(eNPop)
			uncovered := p.st.peek()

			tailNew := p.parserTable.TableReduceCustom(eState, tail)
			p.st.Push(newTokRecChildren(eState, tailNew))

			newState := p.parserTable.TableGoto(uncovered, eState)
			if newState == None {
				return eAct
			}
			p.st.Push(newTokRec(newState))

		case ActAccept:
			return eAct

		case ActNone:
			termNew, termStrNew := p.parserTable.TableRecover(s, term, &p.st, termStr)
			if termNew != TokUndefined {
				term = termNew
				termStr = termStrNew
				continue
			}
			return eAct

		default:
			fmt.Println("Unexpected state >")
			return eAct
		}
	}

	return ActNone
}
