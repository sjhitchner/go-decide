
// generated by gocc; DO NOT EDIT.

package lexer

import (
	
	// "fmt"
	// "github.com/sjhitchner/go-decide/util"
	
	"io/ioutil"
	"unicode/utf8"
	"github.com/sjhitchner/go-decide/token"
)

const(
	NoState = -1
	NumStates = 51
	NumSymbols = 57
)

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, startLine, startColumn, end := this.pos, this.line, this.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {

			switch rune1 {
			case '\n':
				this.line++
				this.column = 1
			case '\r':
				this.column = 1
			case '\t':
				this.column += 4
			default:
				this.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start, startLine, startColumn = this.pos, this.line, this.column
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line ,tok.Pos.Column = start, startLine, startColumn

	
	// fmt.Printf("Token at %s: %s \"%s\"\n", tok.String(), token.TokMap.Id(tok.Type), tok.Lit)
	

	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '.'
1: '_'
2: '.'
3: '"'
4: '"'
5: '''
6: '''
7: 'o'
8: 'r'
9: '>'
10: '>'
11: '='
12: '<'
13: '<'
14: '='
15: '='
16: '='
17: '='
18: '!'
19: '='
20: 'i'
21: 's'
22: 'n'
23: 'o'
24: 't'
25: 'c'
26: 'o'
27: 'n'
28: 't'
29: 'a'
30: 'i'
31: 'n'
32: 's'
33: 'm'
34: 'a'
35: 't'
36: 'c'
37: 'h'
38: 'e'
39: 's'
40: 't'
41: 'r'
42: 'u'
43: 'e'
44: 'f'
45: 'a'
46: 'l'
47: 's'
48: 'e'
49: ' '
50: '\t'
51: '\n'
52: '\r'
53: '0'-'9'
54: 'a'-'z'
55: 'A'-'Z'
56: .

*/