package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror009(t *testing.T){e:=&Error{};if got:=e.ErrorOrNil();got!=nil{t.Fatalf("got=%v",got)}}

func TestGoletaMultierror009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	e:=&Error{Errors:[]error{errors.New("x")}};if got:=e.ErrorOrNil();got==nil{t.Fatal("nonempty returned nil")}
}
