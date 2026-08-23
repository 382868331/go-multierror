package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror012(t *testing.T){target:=errors.New("x");e:=&Error{Errors:[]error{target}};if got:=e.Unwrap();got!=target{t.Fatalf("got=%v",got)}}
