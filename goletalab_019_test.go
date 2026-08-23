package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror019(t *testing.T){e:=&Error{Errors:[]error{errors.New("a"),errors.New("b")}};if e.Len()!=2{t.Fatalf("len=%d",e.Len())}}
