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
		[]float32{1, 2, -3, 4.3345, 1e9},
		nil,
		false,
	}

	// Use the default implementation of GCM
	id := encryption.TokenKeyEncryptionCreatorID("DefaultGCM")

	// For key transfer - envelope encryption details
	gcm, err := encryption.DefaultAlgoFactory.GetAlgorithm(encryption.GCM)
	if err != nil {
		t.Fatalf("envelope - Algorithm retrieval: %v\n", err)
	}
	masterKey, err := gcm.CreateKey()
	if err != nil {
		t.Fatalf("envelope - key generation: %v\n", err)
	}

	for _, i := range testData {

		v, _ := value.NewValue(i)

		for _, encrypt := range []bool{true, false} {

			e, err := encryption.DefaultTokenKeyEncryptionFactory.GetTokenKeyEncryptor(id)
			if err != nil {
				t.Fatalf("TKE Factory - retrieve Encryptor: %v\n", err)
			}

			eo, err := NewEncryptableValue([]byte("Token1"), v, encrypt, e)
			if err != nil {
				t.Fatalf("failed to encrypt: %v\n", err)
			}

			k, err := e.GetKeys(masterKey, gcm)
			if err != nil {
				t.Fatalf("failed to get encryption keys: %v\n", err)
			}

			d, err := encryption.DefaultTokenKeyEncryptionFactory.GetTokenKeyDecryptor(id, masterKey, k, encryption.DefaultAlgoFactory)
			if err != nil {
				t.Fatalf("TKE Factory - retrieve Decryptor: %v\n", err)
			}

			p, _ := NewEncryptableValueParser(d)

			v1, err := p.Parse(eo)
			if err != nil {
				t.Fatalf("failed to decrypt: %v\n", err)
			}

			i1, err := value.ParseValue(v1)
			if err != nil {
				t.Fatalf("failed to parse value: %v\n", err)
			}

			if fmt.Sprint(i1) != fmt.Sprint(i) {
				t.Fatalf("failed to match: wanted %v, got %v\n", i, i1)

			}
		}
	}
}
