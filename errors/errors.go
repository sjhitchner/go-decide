
// generated by gocc; DO NOT EDIT.

package errors

import(
	"bytes"
	"fmt"
	"github.com/sjhitchner/go-decide/token"
)

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
	StackTop       int
}

func (E *Error) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Error")
	if E.Err != nil {
		fmt.Fprintf(w, " %s\n", E.Err)
	} else {
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", E.ErrorToken.Type, E.ErrorToken.Lit)
	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", E.ErrorToken.Pos.Offset, E.ErrorToken.Pos.Line, E.ErrorToken.Pos.Column)
	fmt.Fprintf(w, "Expected one of: ")
	for _, sym := range E.ExpectedTokens {
		fmt.Fprintf(w, "%s ", sym)
	}
	fmt.Fprintf(w, "ErrorSymbol:\n")
	for _, sym := range E.ErrorSymbols {
		fmt.Fprintf(w, "%v\n", sym)
	}
	return w.String()
}

func (e *Error) Error() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Error in S%d: %s, %s", e.StackTop, token.TokMap.TokenString(e.ErrorToken), e.ErrorToken.Pos.String())
	if e.Err != nil {
		w.WriteString(e.Err.Error())
	} else {
		w.WriteString(", expected one of: ")
		for _, expected := range e.ExpectedTokens {
			fmt.Fprintf(w, "%s ", expected)
		}
	}
	return w.String()
}
