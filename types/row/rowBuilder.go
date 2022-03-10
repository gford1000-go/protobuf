package row

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/cell"
	"github.com/gford1000-go/protobuf/types/encrypted_object"
	"github.com/gford1000-go/protobuf/types/encryption"
	"google.golang.org/protobuf/proto"
)

var errMissingEncryptor = errors.New("TokenKeyEncryptor must not be nil")
var errMissingMapper = errors.New("AttributeFromName must not be nil")

func NewRowBuilder(encryptor encryption.TokenKeyEncryptor, mapper AttributeFromName) (*RowBuilder, error) {
	if encryptor == nil {
		return nil, errMissingEncryptor
	}
	if mapper == nil {
		return nil, errMissingMapper
	}
	return &RowBuilder{e: encryptor, m: mapper}, nil
}

type RowBuilder struct {
	e encryption.TokenKeyEncryptor
	m AttributeFromName
}

// Marshal converts the supplied attribute map for the specified row, and
// converts it to an encrypted RowPartition, returned as an EncryptedObject.
//
// The contents of the RowPartition are only retrievable of the requestor
// has access to the key that is specified by the row's rowToken.
func (rb *RowBuilder) Marshal(rowID RowID, i map[AttributeName]interface{}, r RowEncryptionInterfacesRetriever) ([]byte, error) {

	rowToken, err := r.GetRowToken()
	if err != nil {
		return nil, err
	}

	cb, err := cell.NewCellBuilder(rb.e)
	if err != nil {
		return nil, err
	}

	attValues := map[int32][]byte{}

	for k, v := range i {

		attID, err := rb.m.FindIdentifier(k)
		if err != nil {
			return nil, err
		}

		b, err := cb.Marshal(v, r.GetTokenKeyRetriever(k), r.GetValueEncryptionDeterminer())
		if err != nil {
			return nil, err
		}

		attValues[int32(attID)] = b
	}

	rp := &RowPartition{
		RowId:           int64(rowID),
		AttributeValues: attValues,
	}

	eo, err := encrypted_object.NewEncryptedObject(rowToken, rp, rb.e)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(eo)
}
