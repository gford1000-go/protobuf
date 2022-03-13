package value

import (
	"fmt"
	"reflect"
	"strings"
)

func ExampleParseValue() {
	createFloat64 := func(f float64) *float64 {
		p := new(float64)
		*p = f
		return p
	}

	slicePtrToString := func(i interface{}) string {
		v := reflect.ValueOf(i)
		s := ""
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).IsNil() {
				s = s + "<nil> "
			} else {
				s = s + fmt.Sprintf("%v ", v.Index(i).Elem())
			}
		}
		return fmt.Sprintf("[%v]", strings.TrimSpace(s))
	}

	createValueAndParse := func(i interface{}) interface{} {
		v, _ := NewValue(i)
		p, _ := ParseValue(v)
		return p
	}

	i1 := createValueAndParse([]int64{1, 2, 3, 4})
	i2 := createValueAndParse([]*float64{createFloat64(1.99), createFloat64(-2.12), nil})

	fmt.Printf("%v %v, %v %v", i1, reflect.TypeOf(i1), slicePtrToString(i2), reflect.TypeOf(i2))
	// Output: [1 2 3 4] []int64, [1.99 -2.12 <nil>] []*float64
}
