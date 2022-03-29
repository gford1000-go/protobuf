package cell

import (
	"errors"
	"fmt"

	"github.com/gford1000-go/protobuf/types/encryptable_value"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errNonScalarCell = errors.New("non-scalar Cells are not supported")
var errMissingTokenToCell = errors.New("token found in Cell but no CellTokenRetriever provided")

// NewCellParser creates an instance of the parser for a Cell
// A Cell currently only has a Scalar value, which is of type EncryptableValue,
// but in future would support lists and maps to support complex types
func NewCellParser(decryptor encryption.TokenKeyDecryptor, tokenToCell CellTokenRetriever) (*CellParser, error) {
	e, err := encryptable_value.NewEncryptableValueParser(decryptor)
	if err != nil {
		return nil, err
	}

	return &CellParser{e: e, m: tokenToCell}, nil
}

type CellParser struct {
	e *encryptable_value.EncryptableValueParser
	m CellTokenRetriever
}

// Parse returns the scalar Value from a Cell
func (cp *CellParser) Parse(data []byte) (*value.Value, error) {

	// data should be a serialised Cell
	c := &Cell{}
	err := proto.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return cp.parseCell(c)
}

// parseCell unpacks the Value from a Cell, including when the Cell is tokenised
func (cp *CellParser) parseCell(c *Cell) (*value.Value, error) {
	switch x := c.C.(type) {
	case *Cell_Scalar:
		{
			return cp.e.Parse(x.Scalar)
		}
	case *Cell_Token:
		{
			if cp.m == nil {
				return nil, errMissingTokenToCell
			} else {
				if c, err := cp.m.GetCell(x.Token); err != nil {
					return nil, err
				} else {
					return cp.parseCell(c)
				}
			}
		}
	default:
		{
			fmt.Println(x)
			return nil, errNonScalarCell
		}
	}
}
