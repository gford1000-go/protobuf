package row

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/cell"
	"github.com/gford1000-go/protobuf/types/encrypted_object"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errInvalidRowDecryptionInterfacesRetriever = errors.New("RowDecryptionInterfacesRetriever must not be nil")
var errInvalidAttributeMapper = errors.New("AttributeFromIdentifier must not be nil")

// NewRowParser creates an instance of RowParser, which can convert
// RowPartition instances into Rows.
//
// Two interfaces support this activity:
//   1. RowDecryptionInterfacesRetriever provides the TokenKeyDecryptor
//   interfaces for the row and cell decryption
//   2. AttributeFromIdentifier replaces the int64 attribute identifier
//   used in Cell with the attribute name
func NewRowParser(r RowDecryptionInterfacesRetriever, mapper AttributeFromIdentifier) (*RowParser, error) {
	if r == nil {
		return nil, errInvalidRowDecryptionInterfacesRetriever
	}
	p, err := encrypted_object.NewEncryptedObjectParser(r.GetRowTokenKeyDecryptor())
	if err != nil {
		return nil, err
	}
	if mapper == nil {
		return nil, errInvalidAttributeMapper
	}
	return &RowParser{r: r, m: mapper, p: p}, nil
}

type RowParser struct {
	m AttributeFromIdentifier
	r RowDecryptionInterfacesRetriever
	p *encrypted_object.EncryptedObjectParser
}

// Parse returns an instance of Row, unpacking the serialised
// RowPartition data
func (rp *RowParser) Parse(data []byte) (*Row, error) {

	// data should be a marshaled EncryptedObject
	eo := &encrypted_object.EncryptedObject{}
	err := proto.Unmarshal(data, eo)
	if err != nil {
		return nil, err
	}

	// Decrypt into a RowPartition
	row := &RowPartition{}
	err = rp.p.Parse(eo, row)
	if err != nil {
		return nil, err
	}

	newRow := newRow(RowID(row.GetRowId()))

	// Cell parser created for each row, as rows may
	// use different encryptions for their cells
	cp, err := cell.NewCellParser(rp.r.GetCellsTokenKeyDecryptor(newRow.rowId))
	if err != nil {
		return nil, err
	}

	// Unpack each Cell in turn, converting it's Value to
	// a standard Go type
	for attKeyInt, cellAsBytes := range row.GetAttributeValues() {

		attName, err := rp.m.FindName(AttributeIdentifier(attKeyInt))
		if err != nil {
			return nil, err
		}

		v, err := cp.Parse(cellAsBytes)
		if err != nil {
			return nil, err
		}

		attValue, err := value.ParseValue(v)
		if err != nil {
			return nil, err
		}

		err = newRow.addAttribute(attName, attValue)
		if err != nil {
			return nil, err
		}
	}

	return newRow, nil
}
