package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror006(t *testing.T){got:=ListFormatFunc([]error{errors.New("boom")});if !strings.HasPrefix(got,"1 error occurred:"){t.Fatalf("got=%q",got)}}
