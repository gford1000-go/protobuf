package value

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var errTypeMismatch = errors.New("unexpected type mismatch")

// createListFromType returns a strongly typed slice, given the
// parsed values and the specified type
func createListFromType(t ValueListType, l []interface{}) (interface{}, error) {

	switch t {
	case ValueListType_Interface:
		{
			var err error
			var vv *Value
			ll := make([]interface{}, len(l))
			for i, v := range l {
				if reflect.TypeOf(v) == reflect.TypeOf(vv) {
					v, err = ParseValue(v.(*Value))
					if err != nil {
						return nil, err
					}
				}
				ll[i] = v
			}

			return ll, nil
		}
	case ValueListType_Bool:
		{
			ll := make([]bool, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(bool)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrBool:
		{
			ll := make([]*bool, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*bool)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Bytes:
		{
			ll := make([][]byte, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.([]byte)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_Double:
		{
			ll := make([]float64, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(float64)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrDouble:
		{
			ll := make([]*float64, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*float64)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Float:
		{
			ll := make([]float32, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(float32)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrFloat:
		{
			ll := make([]*float32, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*float32)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Int32:
		{
			ll := make([]int32, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(int32)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrInt32:
		{
			ll := make([]*int32, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*int32)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Int64:
		{
			ll := make([]int64, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(int64)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrInt64:
		{
			ll := make([]*int64, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*int64)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_UInt32:
		{
			ll := make([]uint32, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(uint32)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrUInt32:
		{
			ll := make([]*uint32, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*uint32)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_UInt64:
		{
			ll := make([]uint64, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(uint64)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrUInt64:
		{
			ll := make([]*uint64, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*uint64)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_SInt32:
		{
			ll := make([]int32, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(int32)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrSInt32:
		{
			ll := make([]*int32, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*int32)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_SInt64:
		{
			ll := make([]int64, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(int64)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrSInt64:
		{
			ll := make([]*int64, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*int64)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Fixed32:
		{
			ll := make([]uint32, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(uint32)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrFixed32:
		{
			ll := make([]*uint32, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*uint32)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Fixed64:
		{
			ll := make([]uint64, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(uint64)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrFixed64:
		{
			ll := make([]*uint64, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*uint64)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_String:
		{
			ll := make([]string, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(string)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrString:
		{
			ll := make([]*string, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*string)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Time:
		{
			ll := make([]time.Time, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(time.Time)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrTime:
		{
			ll := make([]*time.Time, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*time.Time)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_Duration:
		{
			ll := make([]time.Duration, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(time.Duration)
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_PtrDuration:
		{
			ll := make([]*time.Duration, len(l))
			for i, v := range l {
				if v != nil {
					var ok bool
					ll[i], ok = v.(*time.Duration)
					if !ok {
						return nil, errTypeMismatch
					}
				}
			}

			return ll, nil
		}
	case ValueListType_ValueList:
		{
			ll := make([][]interface{}, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.([]interface{})
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	case ValueListType_ValueMap:
		{
			ll := make([]map[string]interface{}, len(l))
			for i, v := range l {
				var ok bool
				ll[i], ok = v.(map[string]interface{})
				if !ok {
					return nil, errTypeMismatch
				}
			}

			return ll, nil
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("unsupported type in ValueList (%v)", t)
	}
}

// listFromValueList creates a strongly typed list from a ValueList
func listFromValueList(vl *Value_ValueList) (interface{}, error) {

	l := make([]interface{}, len(vl.V))

	for i := 0; i < len(vl.V); i++ {
		ii, err := ParseValue(vl.V[i])
		if err != nil {
			return nil, err
		}
		l[i] = ii
	}

	return createListFromType(vl.T, l)
}

// ParseValue creates an instance of ValueParser that extracts go types
func ParseValue(v *Value) (interface{}, error) {

	if v == nil {
		return nil, nil
	}

	var i interface{}

	switch x := v.V.(type) {
	case *Value_IsNull:
		{
			x = nil
		}
	case *Value_B:
		{
			i = x.B
		}
	case *Value_Pb:
		{
			v := new(bool)
			*v = x.Pb
			i = v
		}
	case *Value_X:
		{
			i = x.X
		}
	case *Value_I32:
		{
			i = x.I32
		}
	case *Value_Pi32:
		{
			v := new(int32)
			*v = x.Pi32
			i = v
		}
	case *Value_I64:
		{
			i = x.I64
		}
	case *Value_Pi64:
		{
			v := new(int64)
			*v = x.Pi64
			i = v
		}
	case *Value_U32:
		{
			i = x.U32
		}
	case *Value_Pu32:
		{
			v := new(uint32)
			*v = x.Pu32
			i = v
		}
	case *Value_U64:
		{
			i = x.U64
		}
	case *Value_Pu64:
		{
			v := new(uint64)
			*v = x.Pu64
			i = v
		}
	case *Value_F32:
		{
			i = x.F32
		}
	case *Value_Pf32:
		{
			v := new(uint32)
			*v = x.Pf32
			i = v
		}
	case *Value_F64:
		{
			i = x.F64
		}
	case *Value_Pf64:
		{
			v := new(uint64)
			*v = x.Pf64
			i = v
		}
	case *Value_F:
		{
			i = x.F
		}
	case *Value_Pf:
		{
			v := new(float32)
			*v = x.Pf
			i = v
		}
	case *Value_D:
		{
			i = x.D
		}
	case *Value_Pd:
		{
			v := new(float64)
			*v = x.Pd
			i = v
		}
	case *Value_S:
		{
			i = x.S
		}
	case *Value_Ps:
		{
			v := new(string)
			*v = x.Ps
			i = v
		}
	case *Value_T:
		{
			i = time.Unix(x.T.Seconds, int64(x.T.Nanos))
		}
	case *Value_Pt:
		{
			v := new(time.Time)
			*v = time.Unix(x.Pt.Seconds, int64(x.Pt.Nanos))
			i = v
		}
	case *Value_Dur:
		{
			i = time.Duration(x.Dur)
		}
	case *Value_Pdur:
		{
			v := new(time.Duration)
			*v = time.Duration(x.Pdur)
			i = v
		}
	case *Value_M:
		{
			m := map[string]interface{}{}

			for k, vv := range x.M.M {

				ii, err := ParseValue(vv)
				if err != nil {
					return nil, err
				}

				m[k] = ii
			}

			i = m
		}
	case *Value_L:
		{
			var err error
			i, err = listFromValueList(x.L)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unsupported type: %v", reflect.TypeOf(x))
	}

	return i, nil
}
