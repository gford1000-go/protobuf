package encryption

// Decryptor will attempt to decrypt the ciphertext using the
// specified key applied the specified algorithm
type Decryptor interface {
	Decrypt(key []byte, a Algorithm, ciphertext []byte) ([]byte, error)
}
