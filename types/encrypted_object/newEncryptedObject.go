package encrypted_object

import (
	"github.com/gford1000-go/protobuf/types/encryption"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// NewEncryptedObject creates an instance of EncryptedObject
// from the supplied message and encryptor details
func NewEncryptedObject(keyToken []byte, message protoreflect.ProtoMessage, encryptor encryption.TokenKeyEncryptor) (*EncryptedObject, error) {

	b, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	e, a, err := encryptor.Encrypt(keyToken, b)
	if err != nil {
		return nil, err
	}

	algo, err := encryption.NewAlgo(a)
	if err != nil {
		return nil, err
	}

	return &EncryptedObject{
			A:        algo,
			KeyToken: keyToken,
			V:        e,
		},
		err
}
