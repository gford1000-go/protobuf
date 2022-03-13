package value

import (
	"bytes"
	"fmt"
	reflect "reflect"
	"testing"
	"time"
)

type testData struct {
	data        interface{}
	expectError bool
}

func testCopyInterface(dst, src interface{}) interface{} {
	if reflect.TypeOf(dst).Kind() == reflect.Ptr {
		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(src))
		return dst
	} else {
		return src
	}
}

func testNewPBool(b bool) *bool {
	p := new(bool)
	*p = b
	return p
}

func testNewPF64(f float64) *float64 {
	p := new(float64)
	*p = f
	return p
}

var testDataSet = []testData{
	{data: "Hello World", expectError: false},
	{data: int32(1234), expectError: false},
	{data: float64(9999.99), expectError: false},
	{data: nil, expectError: false},
	{data: true, expectError: false},
	{data: []interface{}{}, expectError: false},
	{data: map[string]interface{}{}, expectError: false},
	{data: time.Now(), expectError: false},
	{
		data: []interface{}{
			int64(5678),
			float32(-12.34),
			"Lunchtime",
			false,
			nil,
			[]interface{}{},
		},
		expectError: false,
	},
	{
		data: map[string]interface{}{
			"a": "Dinnertime",
			"b": float64(-7e10),
			"c": nil,
			"d": map[string]interface{}{
				"e": []byte("Breakfast"),
			},
			"e": []interface{}{
				false, true,
			},
		},
		expectError: false,
	},
	{
		data:        time.Since(time.Now()),
		expectError: false,
	},
	{data: testCopyInterface(new(bool), true), expectError: false},
	{data: testCopyInterface(new(int32), int32(19)), expectError: false},
	{data: testCopyInterface(new(int64), int64(19)), expectError: false},
	{data: testCopyInterface(new(uint32), uint32(19)), expectError: false},
	{data: testCopyInterface(new(uint64), uint64(19)), expectError: false},
	{data: testCopyInterface(new(float32), float32(19)), expectError: false},
	{data: testCopyInterface(new(float64), float64(19)), expectError: false},
	{data: testCopyInterface(new(string), "Hello World"), expectError: false},
	{data: testCopyInterface(new(time.Time), time.Now()), expectError: false},
	{data: testCopyInterface(new(time.Duration), time.Duration(1234567890)), expectError: false},
	{data: []bool{true, false, false, true}, expectError: false},
	{data: [][]byte{[]byte("Hi"), nil, []byte("There")}, expectError: false},
	{data: []float32{0.1, -2.1}, expectError: false},
	{data: []float64{0.1, -2.1}, expectError: false},
	{data: []int32{1, -3}, expectError: false},
	{data: []int64{1, -3}, expectError: false},
	{data: []uint32{1, 3}, expectError: false},
	{data: []uint64{1, 3}, expectError: false},
	{data: []string{"Hi", "There"}, expectError: false},
	{data: []time.Time{time.Now(), time.Now()}, expectError: false},
	{data: []*bool{testNewPBool(true), nil, testNewPBool(true), testNewPBool(false)}, expectError: false},
	{data: []*float64{testNewPF64(-2.11), nil, testNewPF64(0.008)}, expectError: false},
	{
		data:        &testData{},
		expectError: true,
	},
}

func TestNewValue(t *testing.T) {

	for _, d := range testDataSet {
		_, err := NewValue(d.data)
		if err != nil && !d.expectError {
			t.Errorf("failed to create Value for %v (type: %v)", d.data, reflect.TypeOf(d.data))
		}
		if err == nil && d.expectError {
			t.Errorf("unexpected success in creating Value for %v (type: %v)", d.data, reflect.TypeOf(d.data))
		}
	}
}

func TestParse(t *testing.T) {

	timeTest := func(i interface{}, x time.Time) {
		// time.Time is a special case here, since
		// time.Time.String() is invoked by Sprint
		// but is not meant for debugging only: see
		// https://pkg.go.dev/time#Time.String
		tm, ok := i.(time.Time)
		if !ok {
			t.Fatal("Unexpected error")
		}

		di, _ := x.MarshalText()
		ii, _ := tm.MarshalText()

		if !bytes.Equal(di, ii) {
			t.Errorf("parsed value does not match original for %v (parsed: %v)", di, ii)
		}
	}

	timePtrTest := func(i interface{}, x *time.Time) {
		// time.Time is a special case here, since
		// time.Time.String() is invoked by Sprint
		// but is not meant for debugging only: see
		// https://pkg.go.dev/time#Time.String
		tm, ok := i.(*time.Time)
		if !ok {
			t.Fatal("Unexpected error")
		}

		di, _ := x.MarshalText()
		ii, _ := tm.MarshalText()

		if !bytes.Equal(di, ii) {
			t.Errorf("parsed value does not match original for %v (parsed: %v)", di, ii)
		}
	}

	for _, d := range testDataSet {
		if d.expectError {
			continue
		}
		v, err := NewValue(d.data)
		if err != nil {
			t.Errorf("failed to create Value for %v (type: %v)", d.data, reflect.TypeOf(d.data))
		}

		i, err := ParseValue(v)
		if err != nil {
			t.Errorf("failed to parse Value for %v (type: %v)", d.data, reflect.TypeOf(d.data))
		}

		switch x := d.data.(type) {
		case time.Time:
			{
				timeTest(i, x)
			}
		case *time.Time:
			{
				timePtrTest(i, x)
			}
		default:
			{
				if d.data == nil {
					if i != nil {
						t.Errorf("parsed value does not match original for %v (parsed: %v)", d.data, i)
					}
				} else {
					switch reflect.TypeOf(d.data).Kind() {
					case reflect.Ptr:
						{
							f := func(d interface{}) interface{} {
								return reflect.ValueOf(d).Elem().Interface()
							}

							if fmt.Sprint(f(d.data)) != fmt.Sprint(f(i)) {
								t.Errorf("parsed value does not match original for %v (parsed: %v)", d.data, i)
							}
						}
					case reflect.Slice:
						{
							stringizeSlice := func(o interface{}, sliceType reflect.Type) []string {
								var tt time.Time

								if o == nil {
									t.Fatal("unexpected o == nil")
								}
								v := reflect.ValueOf(o)
								if v == reflect.ValueOf(nil) {
									t.Fatal("unexpected v == nil")
								}
								l := make([]string, v.Len())
								for i := 0; i < v.Len(); i++ {
									if v.Index(i) == reflect.ValueOf(nil) {
										l[i] = ""
									} else {
										if v.Index(i).Type().Kind() == reflect.Ptr && v.Index(i).IsNil() {
											l[i] = "<nil>"
										} else {
											switch sliceType {
											case reflect.TypeOf(tt):
												{
													b, _ := (v.Index(i).Interface().(time.Time)).MarshalText()
													l[i] = string(b)
												}
											case reflect.TypeOf(&tt):
												{
													b, _ := (v.Index(i).Interface().(*time.Time)).MarshalText()
													l[i] = string(b)
												}
											default:
												l[i] = fmt.Sprintf("%v", v.Index(i).String())
											}
										}
									}
								}
								return l
							}

							sliceType := reflect.TypeOf(d.data).Elem()
							if fmt.Sprint(stringizeSlice(d.data, sliceType)) != fmt.Sprint(stringizeSlice(i, sliceType)) {
								t.Errorf("parsed value does not match original for %v (parsed: %v)", stringizeSlice(d.data, sliceType), stringizeSlice(i, sliceType))
							}
						}
					default:
						{
							if fmt.Sprint(d.data) != fmt.Sprint(i) {
								t.Errorf("parsed value does not match original for %v (parsed: %v)", d.data, i)
							}
						}
					}
				}
			}
		}
	}
}
