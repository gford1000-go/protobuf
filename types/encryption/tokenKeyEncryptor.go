package encryption

// // TokenKeyEncryptor receives a token value which is used to retrieve
// // the key required to encrypt.
// //
// // The GetKeys function uses the provided key to encrypt the map
// // of tokens->keys inside an EncryptedObject, for secure distribution
// type TokenKeyEncryptor interface {
// 	EncryptFromToken(token []byte, plaintext []byte) ([]byte, AlgoType, error)
// 	GetKeys(key []byte, a Algorithm) (*EncryptedObject, error)
// }
