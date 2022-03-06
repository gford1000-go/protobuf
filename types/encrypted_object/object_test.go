package encrypted_object

import (
	"fmt"
	"testing"

	"github.com/gford1000-go/protobuf/types/encryption"
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

	e := encryption.NewGCMTokenKeyEncryptor()
	d := encryption.NewGCMTokenKeyDecryptor(e.GetKeys())

	for _, i := range testData {

		v, _ := value.NewValue(i)

		eo, err := NewEncryptedObject([]byte("Token1"), v, e)
		if err != nil {
			t.Errorf("failed to encrypt: %v", err)
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
