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
		[]bool{true, true, false, true},
		map[string]interface{}{
			"a": nil,
			"b": []interface{}{int64(1), true, float32(-0.0000012)},
		},
		false,
	}

	// These are the details of the envelope encryption
	// for keys transfer
	gcm, _ := DefaultAlgoFactory.GetAlgorithm(GCM)
	masterKey, _ := gcm.CreateKey()

	for _, i := range testData {

		id := TokenKeyEncryptionCreatorID("DefaultGCM")

		// Encrypts by key token
		e, _ := DefaultTokenKeyEncryptionFactory.GetTokenKeyEncryptor(id)

		// Dummy value to be serialised
		v, _ := value.NewValue(i)

		// Apply encryption during serialisation
		eo, err := NewEncryptedObjectFromToken([]byte("Token1"), v, e)
		if err != nil {
			t.Errorf("failed to encrypt %v: %v", i, err)
		}

		// Illustrates secure extraction of keys from encryptor,
		// providing the details for their envelope encryption
		encryptedKeys, err := e.GetKeys(masterKey, gcm)
		if err != nil {
			t.Errorf("failed to get encrypted keys: %v", err)
		}

		// Populate using the encrypted keys from the extractor, using
		// the Default AlgorithmFactory to provide the decryption algo
		d, err := DefaultTokenKeyEncryptionFactory.GetTokenKeyDecryptor(id, masterKey, encryptedKeys, DefaultAlgoFactory)
		if err != nil {
			t.Errorf("failed to decrypt keys: %v", err)
		}

		// Create a parser instance that will use the TokenKeyDecryptor
		p, _ := NewEncryptedObjectParser(d)

		// Parse the provided serialised EncryptedObject
		var v1 value.Value
		if err := p.Parse(eo, &v1); err != nil {
			t.Errorf("failed to decrypt: %v", err)
		}

		// Extract out the go instance from the Value
		i1, err := value.ParseValue(&v1)
		if err != nil {
			t.Errorf("failed to parse value: %v", err)
		}

		// Should match
		if fmt.Sprint(i1) != fmt.Sprint(i) {
			t.Errorf("failed to match: wanted %v, got %v", i, i1)
		}
	}

}
