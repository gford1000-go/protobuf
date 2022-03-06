package encryptable_value

import (
	"github.com/gford1000-go/protobuf/types/encrypted_object"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
)

// EncryptableValueParser returns a Value, decrypting if required using
// the supplied TokenKeyDecryptor
func NewEncryptableValueParser(decryptor encryption.TokenKeyDecryptor) (*EncryptableValueParser, error) {
	eop, err := encrypted_object.NewEncryptedObjectParser(decryptor)
	if err != nil {
		return nil, err
	}

	return &EncryptableValueParser{eop: eop}, nil
}

type EncryptableValueParser struct {
	eop *encrypted_object.EncryptedObjectParser
}

// Parse examines the supplied value and extracts the Value from it,
// decrypting if required
func (cp *EncryptableValueParser) Parse(e *EncryptableValue) (*value.Value, error) {

	var v *value.Value

	switch x := e.C.(type) {
	case *EncryptableValue_V:
		{
			v = x.V
		}
	case *EncryptableValue_E:
		{
			var err error
			if err = cp.eop.Parse(x.E, v); err != nil {
				return nil, err
			}
		}
	}

	return v, nil
}
