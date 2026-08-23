package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror003(t *testing.T){base:=errors.New("base");got:=Append(base,errors.New("next"));if len(got.Errors)!=2||got.Errors[0]!=base{t.Fatalf("errors=%v",got.Errors)}}

func TestGoletaMultierror003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	base:=errors.New("only");got:=Append(base);if len(got.Errors)!=1||got.Errors[0]!=base{t.Fatalf("errors=%v",got.Errors)}
}
