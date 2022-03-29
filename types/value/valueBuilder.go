package value

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var errNotSlice = errors.New("not a slice type")
var errUnsupportedSliceType = errors.New("unsupported slice type")

// getTypeOfSlice returns the type of the elements of the slice
func getTypeOfSlice(i interface{}) (reflect.Type, error) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		return reflect.TypeOf(nil), errNotSlice
	}
	return v.Type().Elem(), nil
}

// sliceTypeToValueListTypeMap is the set of supported slice types
var sliceTypeToValueListTypeMap map[reflect.Type]ValueListType = map[reflect.Type]ValueListType{
	reflect.TypeOf([]interface{}{}).Elem():            ValueListType_Interface,
	reflect.TypeOf([]bool{}).Elem():                   ValueListType_Bool,
	reflect.TypeOf([]*bool{}).Elem():                  ValueListType_PtrBool,
	reflect.TypeOf([][]byte{}).Elem():                 ValueListType_Bytes,
	reflect.TypeOf([]float32{}).Elem():                ValueListType_Float,
	reflect.TypeOf([]*float32{}).Elem():               ValueListType_PtrFloat,
	reflect.TypeOf([]float64{}).Elem():                ValueListType_Double,
	reflect.TypeOf([]*float64{}).Elem():               ValueListType_PtrDouble,
	reflect.TypeOf([]int32{}).Elem():                  ValueListType_Int32,
	reflect.TypeOf([]*int32{}).Elem():                 ValueListType_PtrInt32,
	reflect.TypeOf([]int64{}).Elem():                  ValueListType_Int64,
	reflect.TypeOf([]*int64{}).Elem():                 ValueListType_PtrInt64,
	reflect.TypeOf([]uint32{}).Elem():                 ValueListType_UInt32,
	reflect.TypeOf([]*uint32{}).Elem():                ValueListType_PtrUInt32,
	reflect.TypeOf([]uint64{}).Elem():                 ValueListType_UInt64,
	reflect.TypeOf([]*uint64{}).Elem():                ValueListType_PtrUInt64,
	reflect.TypeOf([]string{}).Elem():                 ValueListType_String,
	reflect.TypeOf([]*string{}).Elem():                ValueListType_PtrString,
	reflect.TypeOf([]time.Time{}).Elem():              ValueListType_Time,
	reflect.TypeOf([]*time.Time{}).Elem():             ValueListType_PtrTime,
	reflect.TypeOf([]time.Duration{}).Elem():          ValueListType_Duration,
	reflect.TypeOf([]*time.Duration{}).Elem():         ValueListType_Duration,
	reflect.TypeOf([][]interface{}{}).Elem():          ValueListType_ValueList,
	reflect.TypeOf([]map[string]interface{}{}).Elem(): ValueListType_ValueMap,
}

// fromSliceTypeToValueListType maps from the type of the elements
// of the slice to the ValueListType enumeration value
func fromSliceTypeToValueListType(i interface{}) (ValueListType, error) {

	t, err := getTypeOfSlice(i)
	if err != nil {
		return ValueListType_UnknownValueListType, err
	}

	vlt, ok := sliceTypeToValueListTypeMap[t]
	if !ok {
		return ValueListType_UnknownValueListType, errUnsupportedSliceType
	}

	return vlt, nil
}

// listBuilder creates a Value containing a ValueList
func listBuilder(i interface{}) (*Value, error) {

	t, err := fromSliceTypeToValueListType(i)
	if err != nil {
		return nil, err
	}

	x := reflect.ValueOf(i)

	l := make([]*Value, 0, x.Len())

	for i := 0; i < x.Len(); i++ {
		v, err := NewValue(x.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		l = append(l, v)
	}

	return &Value{
		V: &Value_L{
			L: &Value_ValueList{
				V: l,
				T: t,
			},
		},
	}, nil
}

// NewValue creates an instance of Value holding the specified value
func NewValue(i interface{}) (*Value, error) {

	if i == nil {
		return &Value{V: &Value_IsNull{IsNull: true}}, nil
	}

	var v *Value

	switch x := i.(type) {
	case bool:
		{
			v = &Value{V: &Value_B{B: x}}
		}
	case *bool:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pb{Pb: *x}}
			}
		}
	case []byte:
		{
			v = &Value{V: &Value_X{X: x}}
		}
	case int32:
		{
			v = &Value{V: &Value_I32{I32: x}}
		}
	case *int32:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pi32{Pi32: *x}}
			}
		}
	case int64:
		{
			v = &Value{V: &Value_I64{I64: x}}
		}
	case *int64:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pi64{Pi64: *x}}
			}
		}
	case uint32:
		{
			v = &Value{V: &Value_U32{U32: x}}
		}
	case *uint32:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pu32{Pu32: *x}}
			}
		}
	case uint64:
		{
			v = &Value{V: &Value_U64{U64: x}}
		}
	case *uint64:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pu64{Pu64: *x}}
			}
		}
	case float32:
		{
			v = &Value{V: &Value_F{F: x}}
		}
	case *float32:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pf{Pf: *x}}
			}
		}
	case float64:
		{
			v = &Value{V: &Value_D{D: x}}
		}
	case *float64:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pd{Pd: *x}}
			}
		}
	case string:
		{
			v = &Value{V: &Value_S{S: x}}
		}
	case *string:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Ps{Ps: *x}}
			}
		}
	case time.Time:
		{
			v = &Value{V: &Value_T{
				T: &timestamppb.Timestamp{
					Seconds: x.Unix(),
					Nanos:   int32(x.Nanosecond()),
				}},
			}
		}
	case *time.Time:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pt{
					Pt: &timestamppb.Timestamp{
						Seconds: x.Unix(),
						Nanos:   int32(x.Nanosecond()),
					}},
				}
			}
		}
	case time.Duration:
		{
			v = &Value{V: &Value_Dur{Dur: int64(x)}}
		}
	case *time.Duration:
		{
			if x == nil {
				v = &Value{V: &Value_IsNull{IsNull: true}}
			} else {
				v = &Value{V: &Value_Pdur{Pdur: int64(*x)}}
			}
		}
	case map[string]interface{}:
		{
			m := map[string]*Value{}
			for k, v := range x {
				newV, err := NewValue(v)
				if err != nil {
					return nil, err
				}
				m[k] = newV
			}

			v = &Value{
				V: &Value_M{
					M: &Value_ValueMap{M: m},
				},
			}
		}
	case []interface{},
		[]bool, []*bool, [][]byte,
		[]int64, []*int64,
		[]uint64, []*uint64,
		[]float64, []*float64,
		[]int32, []*int32,
		[]uint32, []*uint32,
		[]float32, []*float32,
		[]string, []*string,
		[]time.Time, []*time.Time,
		[]time.Duration, []*time.Duration,
		[]map[string]interface{}:
		{
			var err error
			v, err = listBuilder(x)
			if err != nil {
				return nil, err
			}
		}
	default:
		{
			vv := reflect.ValueOf(i)
			if vv.Type().Kind() != reflect.Slice {
				return nil, fmt.Errorf("unsupported type: %v", reflect.TypeOf(x))
			}

			l := make([]*Value, vv.Len())
			var err error
			for i := 0; i < vv.Len(); i++ {
				if !vv.Index(i).IsValid() {
					continue
				}
				l[i], err = NewValue(vv.Index(i).Interface())
				if err != nil {
					return nil, err
				}
			}

			v = &Value{
				V: &Value_L{
					L: &Value_ValueList{
						V: l,
						T: ValueListType_Interface,
					},
				},
			}
		}
	}

	return v, nil
}
