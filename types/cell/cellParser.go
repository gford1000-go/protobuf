package cell

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/encryptable_value"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errNonScalarCell = errors.New("non-scalar Cells are not supported")

// NewCellParser creates an instance of the parser for a Cell
// A Cell currently only has a Scalar value, which is of type EncryptableValue,
// but in future would support lists and maps to support complex types
func NewCellParser(decryptor encryption.TokenKeyDecryptor) (*CellParser, error) {
	e, err := encryptable_value.NewEncryptableValueParser(decryptor)
	if err != nil {
		return nil, err
	}

	return &CellParser{e: e}, nil
}

type CellParser struct {
	e *encryptable_value.EncryptableValueParser
}

// Parse returns the scalar Value from a Cell
func (cp *CellParser) Parse(data []byte) (*value.Value, error) {

	c := &Cell{}
	err := proto.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	v := &value.Value{V: &value.Value_IsNull{}}

	switch x := c.C.(type) {
	case *Cell_Scalar:
		{
			v, err = cp.e.Parse(x.Scalar)
		}
	default:
		{
			return nil, errNonScalarCell
		}
	}

	return v, err
}
