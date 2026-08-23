package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror001(t *testing.T){inner:=&Error{Errors:[]error{errors.New("a"),errors.New("b")}};got:=Append(nil,inner);if len(got.Errors)!=2{t.Fatalf("len=%d",len(got.Errors))}}

func TestGoletaMultierror001AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	got:=Append(&Error{},&Error{Errors:[]error{errors.New("x")}});if len(got.Errors)!=1||got.Errors[0].Error()!="x"{t.Fatalf("errors=%v",got.Errors)}
}
