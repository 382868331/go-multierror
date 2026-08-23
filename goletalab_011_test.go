package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror011(t *testing.T){e:=&Error{Errors:[]error{errors.New("a"),errors.New("b")}};got:=e.WrappedErrors();if len(got)!=2||got[0].Error()!="a"{t.Fatalf("got=%v",got)}}

func TestGoletaMultierror011AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	e:=&Error{Errors:[]error{errors.New("x")}};if len(e.WrappedErrors())!=1{t.Fatalf("got=%v",e.WrappedErrors())}
}
