package encryption

import (
	"fmt"
	"testing"

	"github.com/gford1000-go/protobuf/types/value"
)

func TestObject(t *testing.T) {

	testData := []interface{}{
		"Hello World",
		int64(123),
		float32(-17.3),
		nil,
		false,
	}

	masterKey := []byte("0123456789abcdef") // Should be random

	for _, i := range testData {

		e := NewGCMTokenKeyEncryptor()

		v, _ := value.NewValue(i)

		eo, err := NewEncryptedObjectFromToken([]byte("Token1"), v, e)
		if err != nil {
			t.Errorf("failed to encrypt %v: %v", i, err)
		}

		// Illustrates secure extraction of keys from encryptor
		encryptedKeys, err := e.GetKeys(masterKey)
		if err != nil {
			t.Errorf("failed to get encrypted keys: %v", err)
		}

		// Populate using the encrypted keys from the extractor
		d, err := NewGCMTokenKeyDecryptor(masterKey, encryptedKeys)
		if err != nil {
			t.Errorf("failed to decrypt keys: %v", err)
		}

		p, _ := NewEncryptedObjectParser(d)

		var v1 value.Value
		if err := p.Parse(eo, &v1); err != nil {
			t.Errorf("failed to decrypt: %v", err)
		}

		i1, err := value.ParseValue(&v1)
		if err != nil {
			t.Errorf("failed to parse value: %v", err)
		}

		if fmt.Sprint(i1) != fmt.Sprint(i) {
			t.Errorf("failed to match: wanted %v, got %v", i, i1)
		}
	}

}
