package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror012(t *testing.T){target:=errors.New("x");e:=&Error{Errors:[]error{target}};if got:=e.Unwrap();got!=target{t.Fatalf("got=%v",got)}}

func TestGoletaMultierror012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	target:=errors.New("y");e:=&Error{Errors:[]error{target}};if !errors.Is(e,target){t.Fatal("errors.Is failed")}
}
