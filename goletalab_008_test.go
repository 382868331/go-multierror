package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror008(t *testing.T){e:=&Error{Errors:[]error{errors.New("x")},ErrorFormat:func([]error)string{return "custom"}};if got:=e.Error();got!="custom"{t.Fatalf("got=%q",got)}}
