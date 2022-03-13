package value

import (
	"fmt"
	"reflect"
	"time"
)

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
			l := make([]interface{}, 0, len(x.L.V))

			for _, vv := range x.L.V {
				ii, err := ParseValue(vv)
				if err != nil {
					return nil, err
				}
				l = append(l, ii)
			}

			i = l
		}
	default:
		return nil, fmt.Errorf("unsupported type: %v", reflect.TypeOf(x))
	}

	return i, nil
}
