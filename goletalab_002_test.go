package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror002(t *testing.T){got:=Append(nil,errors.New("a"),nil);if len(got.Errors)!=1||got.Errors[0]==nil{t.Fatalf("errors=%v",got.Errors)}}
