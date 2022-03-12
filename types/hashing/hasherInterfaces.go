package hashing

// Hash represents a hash value of something
// Attribute H will always be populated by a Hasher,
// but Alternates is operational, and will depend
// both on the Hasher and the type of thing being hashed.
type Hash struct {
	H          []byte
	Alternates [][]byte
}

// HashType identifies the hashing algorithm used
type HashType uint

const (
	Unknown HashType = iota
	SHA256
)

// Hasher provides a way to generate a Hash
// for the supplied instance.
// Hashers should use a salt within the hashing process
// which is retrievable, to allow consistent hash values
// for the same input
type Hasher interface {
	GetHashType() HashType
	GetSalt() []byte
	Hash(i interface{}) *Hash
}

// HasherCreator can construct instances of a Hasher,
// initialised with the specified salt
type HasherCreator interface {
	New(salt []byte) Hasher
}

// HasherFactory returns a Hasher using the specified
// algorithm, with the ability to specify the salt to be used
type HasherFactory interface {
	AddHasherCreator(c HasherCreator) error
	GetHasher(t HashType) (Hasher, error)
	GetHasherWithSalt(t HashType, salt []byte) (Hasher, error)
}
