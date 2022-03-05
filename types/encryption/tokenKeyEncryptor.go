package encryption

// TokenKeyEncryptor receives a token value which is used to retrieve
// the key required to encrypt.
type TokenKeyEncryptor interface {
	Encrypt(token []byte, plaintext []byte) ([]byte, Algorithm, error)
}
