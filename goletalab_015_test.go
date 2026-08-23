package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror015(t *testing.T){target:=errors.New("first");e:=&Error{Errors:[]error{target,errors.New("last")}};if !errors.Is(e,target){t.Fatal("first target not found")}}
