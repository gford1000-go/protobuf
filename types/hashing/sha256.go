package hashing

import (
	"crypto/sha256"
	"fmt"
)

type sha256Hasher struct {
	s []byte
}

// GetHashType returns the HashType of the algoritm being used
func (s *sha256Hasher) GetHashType() HashType {
	return SHA256
}

// copyBytes is a helper function to copy a byte slice
func (s *sha256Hasher) copyBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

// GetSalt returns the salt used by this instance
func (s *sha256Hasher) GetSalt() []byte {
	return s.copyBytes(s.s)
}

// Hash returns a Hash instance using the algorithm and salt
func (s *sha256Hasher) Hash(i interface{}) *Hash {
	h := sha256.New()
	h.Write(s.s)
	h.Write([]byte(fmt.Sprint(i)))

	return &Hash{
		H: h.Sum(nil),
	}
}

type sha256HasherCreator struct {
}

// New returns a Hasher initialised with the specified salt
func (sc *sha256HasherCreator) New(salt []byte) Hasher {
	s := &sha256Hasher{}
	s.s = s.copyBytes(salt)
	return s
}
