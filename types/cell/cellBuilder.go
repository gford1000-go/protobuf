package cell

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/encryptable_value"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/hashing"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errMissingEncryptor = errors.New("TokenKeyEncryptor must not be nil")
var errMissingHasher = errors.New("Hasher must not be nil")
var errTokenKeyDeterminisorIsNil = errors.New("TokenKeyDeterminisor must not be nil")
var errValueEncryptionDeterminisorIsNil = errors.New("ValueEncryptionDeterminisorIsNil must not be nil")

// NewCellBuilder returns an initialised instance of CellBuilder
func NewCellBuilder(encryptor encryption.TokenKeyEncryptor, hasher hashing.Hasher) (*CellBuilder, error) {
	if encryptor == nil {
		return nil, errMissingEncryptor
	}
	if hasher == nil {
		return nil, errMissingHasher
	}
	return &CellBuilder{e: encryptor, h: hasher}, nil
}

// CellBuilder serialises Cells, optionally applying encryption
type CellBuilder struct {
	e encryption.TokenKeyEncryptor
	h hashing.Hasher
}

// Marshal serialises the specified value, dependent on the value's token and
// whether the token signifies the value should be encrypted or not.
// Marshal also returns a Hash instance representing the value
func (cb *CellBuilder) Marshal(i interface{}, d TokenKeyRetriever, ve ValueEncryptionDeterminer) ([]byte, *hashing.Hash, error) {

	if d == nil {
		return nil, nil, errTokenKeyDeterminisorIsNil
	}

	if ve == nil {
		return nil, nil, errValueEncryptionDeterminisorIsNil
	}

	h := cb.h.Hash(i)

	v, err := value.NewValue(i)
	if err != nil {
		return nil, nil, err
	}

	var keyToken []byte = d.GetToken()
	var encrypt bool = ve.Encrypt(keyToken)

	e, err := encryptable_value.NewEncryptableValue(keyToken, v, encrypt, cb.e)
	if err != nil {
		return nil, nil, err
	}

	c := &Cell{C: &Cell_Scalar{Scalar: e}}

	data, err := proto.Marshal(c)
	if err != nil {
		return nil, nil, err
	}

	return data, h, nil
}
