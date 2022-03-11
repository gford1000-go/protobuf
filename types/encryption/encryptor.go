package encryption

// Encryptor will attempt to use the key to encrypt the plaintext,
// returning the algorithm used as well as the ciphertext
type Encryptor interface {
	Encrypt(key, plaintext []byte) ([]byte, Algorithm, error)
}
