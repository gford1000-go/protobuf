package encryption

// Decryptor will attempt to decrypt the ciphertext using the
// specified key applied the specified AlgoType
type Decryptor interface {
	Decrypt(key []byte, a AlgoType, ciphertext []byte) ([]byte, error)
}
