package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror005(t *testing.T){nested:=&Error{Errors:[]error{errors.New("first"),&Error{Errors:[]error{errors.New("second")}}}};got:=Flatten(nested).(*Error);if len(got.Errors)!=2||got.Errors[0].Error()!="first"{t.Fatalf("errors=%v",got.Errors)}}

func TestGoletaMultierror005AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	nested:=&Error{Errors:[]error{errors.New("x")}};got:=Flatten(nested).(*Error);if len(got.Errors)!=1{t.Fatalf("errors=%v",got.Errors)}
}
