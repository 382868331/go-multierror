package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror018(t *testing.T){e:=&Error{Errors:[]error{errors.New("a"),errors.New("b")}};got:=Prefix(e,"ctx:").(*Error);if got.Errors[0].Error()!="ctx: a"||got.Errors[1].Error()!="ctx: b"{t.Fatalf("got=%v",got.Errors)}}
