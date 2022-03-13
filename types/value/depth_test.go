package value

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDepth(t *testing.T) {

	data := []interface{}{
		[][]int64{{1, 2}, {3, 4}},
		[][][]int64{{{1, 2}, {3, 4}}, {{5, 6}}},
		[][][]string{
			{
				{
					"Hello",
					"World",
				},
				{
					"Bonjour",
				},
			},
			{
				{
					"Good",
					"Night",
				},
			},
		},
		[]map[string]interface{}{
			{
				"a": [][]int64{{1, 2}, {3, 4}},
				"b": map[string]interface{}{
					"x": map[string]interface{}{
						"m": []int64{1, 2, 3, 4},
						"n": []int64{19, 22, 33, 45},
					},
					"y": []map[string]interface{}{
						{
							"m": []int64{0, -1, 3, 17},
							"n": []int64{22, 0, 88, 2},
						},
					},
				},
				"c": true,
			},
			{},
			{
				"d": "Hello",
			},
		},
		[][]interface{}{
			{
				int64(2),
				"Hello",
				false,
			},
			{},
			{
				[]interface{}{},
				map[string]interface{}{},
				[]interface{}{int32(2)},
				[]float64{1.23, 3.45},
			},
		},
	}

	for _, d := range data {
		v, err := NewValue(d)
		if err != nil {
			t.Fatalf("failed to create Value for type %v (%v) - %v", reflect.TypeOf(d), d, err)
		}

		p, err := ParseValue(v)
		if err != nil {
			t.Fatalf("failed to parse Value for type %v (%v) - %v", reflect.TypeOf(d), d, err)
		}

		if fmt.Sprint(d) != fmt.Sprint(p) {
			t.Fatalf("incorrect parsing for type %v (%v) - %v", reflect.TypeOf(d), d, err)
		}
	}

}
