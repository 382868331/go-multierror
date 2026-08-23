package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror010(t *testing.T){e:=&Error{Errors:[]error{errors.New("x")}};if got:=e.GoString();!strings.HasPrefix(got,"*multierror.Error"){t.Fatalf("got=%q",got)}}

func TestGoletaMultierror010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	e:=&Error{};if got:=e.GoString();len(got)==0||got[0]!='*'{t.Fatalf("got=%q",got)}
}
