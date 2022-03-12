package hashing

import (
	"crypto/rand"
	"errors"
	"io"
)

var errInvalidHasherCreator = errors.New("HasherCreator must not be nil")
var errCreatorDoesNotEmitHasher = errors.New("HasherCreator fails to emit Hashers")
var errSaltGnerationFailure = errors.New("failed to create random salt")
var errUnknownHashType = errors.New("unknown HashType requested")

func init() {
	DefaultFactory, _ = NewFactory([]HasherCreator{
		&sha256HasherCreator{},
	})
}

// DefaultFactory is a Factory pre-filled with existing HashTypes,
// currently SHA256
var DefaultFactory *Factory

// NewFactory returns an instance of Factory, pre-filled with
// the specified set of HasherCreators
func NewFactory(cs []HasherCreator) (*Factory, error) {
	f := &Factory{
		m: make(map[HashType]HasherCreator),
	}

	for _, c := range cs {
		if err := f.AddHasherCreator(c); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// Factory manufactures instances of Hasher by invoking the
// HasherCreator for the required HashType
type Factory struct {
	m map[HashType]HasherCreator
}

// AddHasherCreator inserts the specified HasherCreator
// into the Factory.
// The HashType of the Hashers that the HasherCreator returns
// will be set to this HasherCreator, overwriting the prior
// HasherCreator if that existed.
func (f *Factory) AddHasherCreator(c HasherCreator) error {
	if c == nil {
		return errInvalidHasherCreator
	}

	h := c.New([]byte(""))
	if h == nil {
		return errCreatorDoesNotEmitHasher
	}

	// Will replace existing for the same HashType
	f.m[h.GetHashType()] = c

	return nil
}

// GetHasher returns an instance of a Hasher of the specified
// HashType, initialised with a random salt
func (f *Factory) GetHasher(t HashType) (Hasher, error) {
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, errSaltGnerationFailure
	}
	return f.GetHasherWithSalt(t, salt)
}

// GetHasherWithSalt returns an instance of a Hasher of the specified
// HashType, initialised with the specified salt
func (f *Factory) GetHasherWithSalt(t HashType, salt []byte) (Hasher, error) {
	c, ok := f.m[t]
	if !ok {
		return nil, errUnknownHashType
	}
	return c.New(salt), nil
}
