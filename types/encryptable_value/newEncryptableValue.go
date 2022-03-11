package encryptable_value

import (
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
)

// NewEncryptableValue creats an EncryptableValue from a Value
func NewEncryptableValue(keyToken []byte, v *value.Value, encryptionRequired bool, encryptor encryption.TokenKeyEncryptor) (*EncryptableValue, error) {

	e := &EncryptableValue{}

	if encryptionRequired {
		eo, err := encryption.NewEncryptedObjectFromToken(keyToken, v, encryptor)
		if err != nil {
			return nil, err
		}
		e.C = &EncryptableValue_E{E: eo}
	} else {
		e.C = &EncryptableValue_V{V: v}
	}

	return e, nil
}
