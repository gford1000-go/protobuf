package value

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func ExampleNewValue() {

	createAndGetMarshalledLength := func(i interface{}) int {
		v, _ := NewValue(i)
		b, _ := proto.Marshal(v)
		return len(b)
	}

	data := []interface{}{
		"Ultimate Question of Life, the Universe, and Everything",
		uint64(42),
		float64(1e80),
		map[string]interface{}{
			"q": "Ultimate Question of Life, the Universe, and Everything",
			"a": int64(42),
			"n": float64(1e80),
		},
		[]int32{1, 2, 3},
		[]interface{}{int32(1), int32(2), int32(3)},
	}
	lengths := make([]int, len(data))

	for i, d := range data {
		lengths[i] = createAndGetMarshalledLength(d)
	}

	fmt.Println(lengths)
	// Output: [57 3 9 92 17 18]
}
