package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror013(t *testing.T){a:=errors.New("a");b:=errors.New("b");c:=errors.New("c");e:=&Error{Errors:[]error{a,b,c}};if !errors.Is(e,b){t.Fatal("second error unreachable")}}
