package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror005(t *testing.T){nested:=&Error{Errors:[]error{errors.New("first"),&Error{Errors:[]error{errors.New("second")}}}};got:=Flatten(nested).(*Error);if len(got.Errors)!=2||got.Errors[0].Error()!="first"{t.Fatalf("errors=%v",got.Errors)}}
