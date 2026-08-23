package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror007(t *testing.T){got:=ListFormatFunc([]error{errors.New("a"),errors.New("b")});if !strings.Contains(got,"* a")||!strings.Contains(got,"* b"){t.Fatalf("got=%q",got)}}
