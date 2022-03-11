package encryption

// TokenKeyDecryptor receives a token value and algorithm, which is used to retrieve
// the key required to decrypt and the algo to use to attempt decryption.
// Behaviour when the key is not available is unspecified.
type TokenKeyDecryptor interface {
	DecryptFromToken(token []byte, a Algorithm, ciphertext []byte) ([]byte, error)
}
