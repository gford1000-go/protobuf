package encryption

import (
	"errors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var errMissingDecryptor = errors.New("TokenKeyDecryptor must not be nil")
var errNilMessage = errors.New("message must not be nil")

// EncryptedObjectParser decrypts EncryptedObjects in a Message, using
// the supplied TokenKeyDecryptor
func NewEncryptedObjectParser(decryptor TokenKeyDecryptor) (*EncryptedObjectParser, error) {
	if decryptor == nil {
		return nil, errMissingDecryptor
	}

	return &EncryptedObjectParser{d: decryptor}, nil
}

type EncryptedObjectParser struct {
	d TokenKeyDecryptor
}

// Parse decrypts using into the supplied ProtoMessage instance
func (cp *EncryptedObjectParser) Parse(e *EncryptedObject, message protoreflect.ProtoMessage) error {

	if message == nil {
		return errNilMessage
	}

	a, err := ParseAlgo(e.A)
	if err != nil {
		return err
	}

	b, err := cp.d.DecryptFromToken(e.GetKeyToken(), a, e.V)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(b, message)
	if err != nil {
		return err
	}

	return nil
}
