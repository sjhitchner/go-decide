
// generated by gocc; DO NOT EDIT.

package parser

import ( "github.com/sjhitchner/go-decide/expression" )

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Expression	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Term "or" Term	<< expression.NewLogicalOr(X[0], X[2]) >>`,
		Id: "Expression",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewLogicalOr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expression : Term	<<  >>`,
		Id: "Expression",
		NTType: 1,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor ">" Factor	<< expression.NewComparisonGreaterThan(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonGreaterThan(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor ">=" Factor	<< expression.NewComparisonGreaterThanEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonGreaterThanEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "<" Factor	<< expression.NewComparisonLessThan(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonLessThan(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "<=" Factor	<< expression.NewComparisonLessThanEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonLessThanEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "=" Factor	<< expression.NewComparisonEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "==" Factor	<< expression.NewComparisonEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "!=" Factor	<< expression.NewComparisonNotEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonNotEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "is" Factor	<< expression.NewComparisonIs(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonIs(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "is" "not" Factor	<< expression.NewComparisonIsNot(X[0], X[3]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 11,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonIsNot(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "contains" Factor	<< expression.NewComparisonContains(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewComparisonContains(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "matches" Factor	<< expression.NewMatches(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewMatches(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id: "Term",
		NTType: 2,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : variable	<< expression.NewResolver(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewResolver(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : string_lit	<< expression.NewLiteralString(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewLiteralString(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : int_lit	<< expression.NewLiteralInt(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewLiteralInt(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : float_lit	<< expression.NewLiteralFloat(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewLiteralFloat(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : "true"	<< expression.NewLiteralBool(true) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewLiteralBool(true)
		},
	},
	ProdTabEntry{
		String: `Factor : "false"	<< expression.NewLiteralBool(false) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expression.NewLiteralBool(false)
		},
	},
	
}
