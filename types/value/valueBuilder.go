package value

import (
	"fmt"
	"reflect"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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
			v = &Value{V: &Value_Pb{Pb: *x}}
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
			v = &Value{V: &Value_Pi32{Pi32: *x}}
		}
	case int64:
		{
			v = &Value{V: &Value_I64{I64: x}}
		}
	case *int64:
		{
			v = &Value{V: &Value_Pi64{Pi64: *x}}
		}
	case uint32:
		{
			v = &Value{V: &Value_U32{U32: x}}
		}
	case *uint32:
		{
			v = &Value{V: &Value_Pu32{Pu32: *x}}
		}
	case uint64:
		{
			v = &Value{V: &Value_U64{U64: x}}
		}
	case *uint64:
		{
			v = &Value{V: &Value_Pu64{Pu64: *x}}
		}
	case float32:
		{
			v = &Value{V: &Value_F{F: x}}
		}
	case *float32:
		{
			v = &Value{V: &Value_Pf{Pf: *x}}
		}
	case float64:
		{
			v = &Value{V: &Value_D{D: x}}
		}
	case *float64:
		{
			v = &Value{V: &Value_Pd{Pd: *x}}
		}
	case string:
		{
			v = &Value{V: &Value_S{S: x}}
		}
	case *string:
		{
			v = &Value{V: &Value_Ps{Ps: *x}}
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
			v = &Value{V: &Value_Pt{
				Pt: &timestamppb.Timestamp{
					Seconds: x.Unix(),
					Nanos:   int32(x.Nanosecond()),
				}},
			}
		}
	case time.Duration:
		{
			v = &Value{V: &Value_Dur{Dur: int64(x)}}
		}
	case *time.Duration:
		{
			v = &Value{V: &Value_Pdur{Pdur: int64(*x)}}
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
	case []interface{}:
		{
			l := make([]*Value, 0, len(x))

			for _, v := range x {
				newV, err := NewValue(v)
				if err != nil {
					return nil, err
				}
				l = append(l, newV)
			}

			v = &Value{
				V: &Value_L{
					L: &Value_ValueList{V: l},
				},
			}
		}
	default:
		return nil, fmt.Errorf("unsupported type: %v", reflect.TypeOf(x))
	}

	return v, nil
}
