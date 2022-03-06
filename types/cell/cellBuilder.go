package cell

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/encryptable_value"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errMissingEncryptor = errors.New("TokenKeyEncryptor must not be nil")
var errTokenKeyDeterminisorIsNil = errors.New("TokenKeyDeterminisor must not be nil")
var errValueEncryptionDeterminisorIsNil = errors.New("ValueEncryptionDeterminisorIsNil must not be nil")

func NewCellBuilder(encryptor encryption.TokenKeyEncryptor) (*CellBuilder, error) {
	if encryptor == nil {
		return nil, errMissingEncryptor
	}
	return &CellBuilder{e: encryptor}, nil
}

type CellBuilder struct {
	e encryption.TokenKeyEncryptor
}

func (cb *CellBuilder) Marshal(i interface{}, d TokenKeyRetriever, ve ValueEncryptionDeterminer) ([]byte, error) {

	if d == nil {
		return nil, errTokenKeyDeterminisorIsNil
	}

	if ve == nil {
		return nil, errValueEncryptionDeterminisorIsNil
	}

	v, err := value.NewValue(i)
	if err != nil {
		return nil, err
	}

	var keyToken []byte = d.GetToken()
	var encrypt bool = ve.Encrypt(keyToken)

	e, err := encryptable_value.NewEncryptableValue(keyToken, v, encrypt, cb.e)
	if err != nil {
		return nil, err
	}

	c := &Cell{C: &Cell_Scalar{Scalar: e}}

	data, err := proto.Marshal(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}
