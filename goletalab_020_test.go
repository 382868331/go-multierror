package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror020(t *testing.T){e:=&Error{Errors:[]error{errors.New("b"),errors.New("a")}};sort.Sort(e);if e.Errors[0].Error()!="a"{t.Fatalf("errors=%v",e.Errors)}}
