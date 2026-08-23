package multierror
import("errors";"fmt";"sort";"strings";"testing")
var _=errors.New;var _=fmt.Sprint;var _=sort.Sort;var _=strings.Contains
type typed015 struct{s string}
func(e *typed015)Error()string{return e.s}
func TestGoletaMultierror014(t *testing.T){first:=&typed015{"first"};e:=&Error{Errors:[]error{first,errors.New("last")}};var got *typed015;if !errors.As(e,&got)||got!=first{t.Fatalf("got=%v",got)}}

func TestGoletaMultierror014AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	last:=&typed015{"last"};e:=&Error{Errors:[]error{errors.New("first"),last}};var got *typed015;if !errors.As(e,&got)||got!=last{t.Fatalf("got=%v",got)}
}
