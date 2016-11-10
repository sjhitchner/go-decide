
// generated by gocc; DO NOT EDIT.

package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"or",
		"not",
		">",
		">=",
		"<",
		"<=",
		"=",
		"==",
		"!=",
		"is",
		"contains",
		"matches",
		"variable",
		"string_lit",
		"int_lit",
		"float_lit",
		"true",
		"false",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"or": 2,
		"not": 3,
		">": 4,
		">=": 5,
		"<": 6,
		"<=": 7,
		"=": 8,
		"==": 9,
		"!=": 10,
		"is": 11,
		"contains": 12,
		"matches": 13,
		"variable": 14,
		"string_lit": 15,
		"int_lit": 16,
		"float_lit": 17,
		"true": 18,
		"false": 19,
	},
}

