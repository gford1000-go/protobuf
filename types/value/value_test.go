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

var data = []testData{
	{data: "Hello World", expectError: false},
	{data: int32(1234), expectError: false},
	{data: float64(9999.99), expectError: false},
	{data: nil, expectError: false},
	{data: true, expectError: false},
	{data: []interface{}{}, expectError: false},
	{data: map[string]interface{}{}, expectError: false},
	{data: time.Now(), expectError: false},
	{data: new(int64), expectError: false},
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
		expectError: true,
	},
	{
		data:        &testData{},
		expectError: true,
	},
}

func TestNewValue(t *testing.T) {
	for _, d := range data {
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
	for _, d := range data {
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
