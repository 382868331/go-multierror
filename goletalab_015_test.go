package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror015(t *testing.T){target:=errors.New("first");e:=&Error{Errors:[]error{target,errors.New("last")}};if !errors.Is(e,target){t.Fatal("first target not found")}}

func TestGoletaMultierror015AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	target:=errors.New("last");e:=&Error{Errors:[]error{errors.New("first"),target}};if !errors.Is(e,target){t.Fatal("last target not found")}
}
