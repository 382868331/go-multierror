package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror004(t *testing.T){e:=errors.New("plain");if got:=Flatten(e);got!=e{t.Fatalf("got=%T %v",got,got)}}
