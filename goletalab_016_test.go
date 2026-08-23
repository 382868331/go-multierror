package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror016(t *testing.T){if got:=Prefix(nil,"scope:");got!=nil{t.Fatalf("got=%v",got)}}
