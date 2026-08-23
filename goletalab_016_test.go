package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror016(t *testing.T){if got:=Prefix(nil,"scope:");got!=nil{t.Fatalf("got=%v",got)}}

func TestGoletaMultierror016AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	var e *Error;if got:=Prefix(e,"scope:");got==nil{t.Fatal("typed nil should become Error")}
}
