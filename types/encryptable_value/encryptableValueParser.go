package encryptable_value

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/encrypted_object"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errMissingDecryptor = errors.New("TokenKeyDecryptor must not be nil")
var errUnknownAlgorithmUsed = errors.New("unsupported algorithm used for encryption")

// EncryptableValueParser returns a Value, decrypting if required using
// the supplied TokenKeyDecryptor
func NewEncryptableValueParser(decryptor encryption.TokenKeyDecryptor) (*EncryptableValueParser, error) {
	if decryptor == nil {
		return nil, errMissingDecryptor
	}

	return &EncryptableValueParser{d: decryptor}, nil
}

type EncryptableValueParser struct {
	d encryption.TokenKeyDecryptor
}

func (cp *EncryptableValueParser) decryptValue(e *encrypted_object.EncryptedObject) (*value.Value, error) {

	a, err := encryption.ParseAlgo(e.A)
	if err != nil {
		return nil, err
	}

	b, err := cp.d.Decrypt(e.GetKeyToken(), a, e.V)
	if err != nil {
		return nil, err
	}

	v := &value.Value{}
	err = proto.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}

	return v, nil
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
			v, err = cp.decryptValue(x.E)
			if err != nil {
				return nil, err
			}
		}
	}

	return v, nil
}
