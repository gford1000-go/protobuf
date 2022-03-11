package encryption

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// NewEncryptedObjectFromToken creates an instance of EncryptedObject
// from the supplied message and encryptor details
func NewEncryptedObjectFromToken(keyToken []byte, message protoreflect.ProtoMessage, encryptor TokenKeyEncryptor) (*EncryptedObject, error) {

	b, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	e, a, err := encryptor.EncryptFromToken(keyToken, b)
	if err != nil {
		return nil, err
	}

	algo, err := NewAlgo(a)
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

// NewEncryptedObject creates an instance of EncryptedObject
// from the supplied message and encryptor details
func NewEncryptedObject(key []byte, message protoreflect.ProtoMessage, encryptor Encryptor) (*EncryptedObject, error) {

	b, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	e, a, err := encryptor.Encrypt(key, b)
	if err != nil {
		return nil, err
	}

	algo, err := NewAlgo(a)
	if err != nil {
		return nil, err
	}

	return &EncryptedObject{
			A:        algo,
			KeyToken: nil,
			V:        e,
		},
		err
}
