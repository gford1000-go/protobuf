package encryptable_value

import (
	"fmt"
	"testing"

	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
)

func TestValue(t *testing.T) {

	testData := []interface{}{
		"Hello World",
		int64(123),
		float32(-17.3),
		nil,
		false,
	}

	masterKey := []byte("0123456789abcdef") // Should be random

	for _, i := range testData {

		v, _ := value.NewValue(i)

		for _, encrypt := range []bool{true, false} {

			e := encryption.NewGCMTokenKeyEncryptor()

			eo, err := NewEncryptableValue([]byte("Token1"), v, encrypt, e)
			if err != nil {
				t.Errorf("failed to encrypt: %v", err)
			}

			k, err := e.GetKeys(masterKey)
			if err != nil {
				t.Errorf("failed to get encryption keys: %v", err)
			}

			d, err := encryption.NewGCMTokenKeyDecryptor(masterKey, k)
			if err != nil {
				t.Errorf("failed to create decryptor: %v", err)
			}

			p, _ := NewEncryptableValueParser(d)

			v1, err := p.Parse(eo)
			if err != nil {
				t.Errorf("failed to decrypt: %v", err)
			}

			i1, err := value.ParseValue(v1)
			if err != nil {
				t.Errorf("failed to parse value: %v", err)
			}

			if fmt.Sprint(i1) != fmt.Sprint(i) {
				t.Errorf("failed to match: wanted %v, got %v", i, i1)

			}
		}
	}
}
