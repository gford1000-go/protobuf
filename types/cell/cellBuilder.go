package cell

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/encryptable_value"
	"github.com/gford1000-go/protobuf/types/encrypted_object"
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

func (cb *CellBuilder) marshalValue(v *value.Value) ([]byte, error) {
	data, err := proto.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cb *CellBuilder) encryptValue(keyToken []byte, v *value.Value) (*encrypted_object.EncryptedObject, error) {

	b, err := cb.marshalValue(v)
	if err != nil {
		return nil, err
	}

	e, a, err := cb.e.Encrypt(keyToken, b)
	if err != nil {
		return nil, err
	}

	algo, err := encryption.NewAlgo(a)
	if err != nil {
		return nil, err
	}

	return &encrypted_object.EncryptedObject{
			A:        algo,
			KeyToken: keyToken,
			V:        e,
		},
		err
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

	var e *encrypted_object.EncryptedObject

	if encrypt {
		var err error
		e, err = cb.encryptValue(keyToken, v)
		if err != nil {
			return nil, err
		}
	}

	var c *Cell
	if encrypt {
		c = &Cell{C: &Cell_Scalar{
			Scalar: &encryptable_value.EncryptableValue{
				C: &encryptable_value.EncryptableValue_E{E: e},
			},
		},
		}
	} else {
		c = &Cell{C: &Cell_Scalar{
			Scalar: &encryptable_value.EncryptableValue{
				C: &encryptable_value.EncryptableValue_V{V: v},
			},
		},
		}
	}

	data, err := proto.Marshal(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}
