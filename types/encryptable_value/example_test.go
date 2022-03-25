package encryptable_value

import (
	"fmt"

	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/value"
)

func ExampleNewEncryptableValue() {

	// Value to be serialised
	v, _ := value.NewValue("Hello World")

	// Encryption identifier and factory
	id := encryption.TokenKeyEncryptionCreatorID("DefaultGCM")
	factory := encryption.DefaultTokenKeyEncryptionFactory

	// Get encryptor from factory
	encryptor, _ := factory.GetTokenKeyEncryptor(id)

	// Create encrypted, serialised value, associated with token
	hwToken := []byte("HelloWorldToken")
	eo, _ := NewEncryptableValue(hwToken, v, true, encryptor)

	// To deserialise - first create decryptor by securely
	// transferring keys from the encryptor
	algoFactory := encryption.DefaultAlgoFactory
	gcm, _ := algoFactory.GetAlgorithm(encryption.GCM)
	masterKey, _ := gcm.CreateKey()
	keys, _ := encryptor.GetKeys(masterKey, gcm)
	decryptor, _ := factory.GetTokenKeyDecryptor(id, masterKey, keys, algoFactory)

	// Then create a parser
	p, _ := NewEncryptableValueParser(decryptor)

	// And extract the parsed Value
	v1, _ := p.Parse(eo)

	// Let's retrieve the interface{} in the parsed Value
	i1, _ := value.ParseValue(v1)

	// And show it contains the correct value
	fmt.Printf("Parsed '%v'\n", i1)
	// Output: Parsed 'Hello World'
}
