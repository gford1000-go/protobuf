package encryption

// AlgoType identifies the encryption algorithm used
type AlgoType string

const (
	Unknown AlgoType = "Unknown"
	GCM     AlgoType = "GCM"
)

// Algorithm provides an Encryptor and Decryptor interface,
// that implement the specified AlgoType
type Algorithm interface {
	GetType() AlgoType
	GetEncryptor() Encryptor
	GetDecryptor() Decryptor
}

// AlgorithmCreator can construct instances of a Algorithm
type AlgorithmCreator interface {
	New() Algorithm
}

// AlgorithmFactory returns a Algorithm using the specified
// algorithm
type AlgorithmFactory interface {
	AddAlgorithmCreator(a AlgorithmCreator) error
	GetAlgorithm(t AlgoType) (Algorithm, error)
}

// Decryptor will attempt to decrypt the ciphertext
// using the specified key
type Decryptor interface {
	Decrypt(key []byte, ciphertext []byte) ([]byte, error)
}

// Encryptor will attempt to use the key to encrypt the plaintext,
// returning the AlgoType used as well as the ciphertext
type Encryptor interface {
	Encrypt(key, plaintext []byte) ([]byte, AlgoType, error)
}
