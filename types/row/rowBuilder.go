package row

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/cell"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/hashing"
	"google.golang.org/protobuf/proto"
)

type AttributeHashMap map[AttributeName]*hashing.Hash

var errMissingHasher = errors.New("AttributeHashRetriever must not be nil")
var errMissingEncryptor = errors.New("TokenKeyEncryptor must not be nil")
var errMissingMapper = errors.New("AttributeFromName must not be nil")

func NewRowBuilder(hashRetriever AttributeHashRetriever, encryptor encryption.TokenKeyEncryptor, mapper AttributeFromName) (*RowBuilder, error) {
	if hashRetriever == nil {
		return nil, errMissingHasher
	}
	if encryptor == nil {
		return nil, errMissingEncryptor
	}
	if mapper == nil {
		return nil, errMissingMapper
	}
	return &RowBuilder{e: encryptor, r: hashRetriever, m: mapper}, nil
}

type RowBuilder struct {
	e encryption.TokenKeyEncryptor
	m AttributeFromName
	r AttributeHashRetriever
}

// Marshal converts the supplied attribute map for the specified row, and
// converts it to an encrypted RowPartition, returned as an EncryptedObject.
//
// The contents of the RowPartition are only retrievable of the requestor
// has access to the key that is specified by the row's rowToken.
//
// The AttributeHashMap provides the hash values for each of the attributes
func (rb *RowBuilder) Marshal(rowID RowID, i map[AttributeName]interface{}, r RowEncryptionInterfacesRetriever) ([]byte, AttributeHashMap, error) {

	rowToken, err := r.GetRowToken()
	if err != nil {
		return nil, nil, err
	}

	cb, err := cell.NewCellBuilder(rb.e)
	if err != nil {
		return nil, nil, err
	}

	attValues := map[int32][]byte{}
	attHashMap := AttributeHashMap{}

	for k, v := range i {

		attID, err := rb.m.FindIdentifier(k)
		if err != nil {
			return nil, nil, err
		}

		// Retrieve the hasher for the attribute. This is invariant across rows
		// so that all values of an attribute are hashed the same way, providing
		// the basis of a search mechanism based on the hashes.
		// This also allows different Hashers to be applied for each attribute, if desired.
		hasher, err := rb.r.GetHasher(k)
		if err != nil {
			return nil, nil, err
		}

		b, h, err := cb.Marshal(v, hasher, r.GetTokenKeyRetriever(k), r.GetValueEncryptionDeterminer())
		if err != nil {
			return nil, nil, err
		}

		attValues[int32(attID)] = b
		attHashMap[k] = h
	}

	rp := &RowPartition{
		RowId:           int64(rowID),
		AttributeValues: attValues,
	}

	// Apply consistent row level encryption - so that all rows with the same
	// rowToken are encrypted with the same key and algorithm.  This effectively
	// acts as a filter where consumers do not have access to the key for the rowToken
	eo, err := encryption.NewEncryptedObjectFromToken(rowToken, rp, rb.e)
	if err != nil {
		return nil, nil, err
	}

	b, err := proto.Marshal(eo)
	if err != nil {
		return nil, nil, err
	}

	return b, attHashMap, nil
}
