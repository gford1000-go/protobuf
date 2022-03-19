package encryption

import (
	"errors"
	"sort"
	"sync"
)

// TokenKeyDecryptor receives a token value and AlgoType, which is used to retrieve
// the key required to decrypt and the algo to use to attempt decryption.
// Behaviour when the key is not available is unspecified.
type TokenKeyDecryptor interface {
	DecryptFromToken(token []byte, a AlgoType, ciphertext []byte) ([]byte, error)
}

// TokenKeyEncryptor receives a token value which is used to retrieve
// the key required to encrypt.
//
// The GetKeys function uses the provided key to encrypt the map
// of tokens->keys inside an EncryptedObject, for secure distribution
type TokenKeyEncryptor interface {
	EncryptFromToken(token []byte, plaintext []byte) ([]byte, AlgoType, error)
	GetKeys(key []byte, a Algorithm) (*EncryptedObject, error)
}

// TokenKeyEncryptionFactory returns the encryptor or decryptor
// for the specified TokenKeyEncryptionCreatorID
type TokenKeyEncryptionFactory interface {
	GetTokenKeyEncryptionCreatorIDs() TokenKeyEncryptionCreatorIDList
	AddTokenKeyEncryptionCreator(c TokenKeyEncryptionCreator) error
	GetTokenKeyDecryptor(i TokenKeyEncryptionCreatorID, key []byte, keys *EncryptedObject, factory AlgorithmFactory) (TokenKeyDecryptor, error)
	GetTokenKeyEncryptor(i TokenKeyEncryptionCreatorID) (TokenKeyEncryptor, error)
}

// TokenKeyEncryptionCreatorID identifies TokenKeyEncryptionCreators
type TokenKeyEncryptionCreatorID string

// TokenKeyEncryptionCreator can manufacture encryptors and decryptors
type TokenKeyEncryptionCreator interface {
	GetID() TokenKeyEncryptionCreatorID
	GetEncryptionAlgoType() AlgoType
	GetTokenKeyDecryptor(key []byte, keys *EncryptedObject, factory AlgorithmFactory) (TokenKeyDecryptor, error)
	GetTokenKeyEncryptor() (TokenKeyEncryptor, error)
}

// TokenKeyEncryptionCreatorIDList is a slice of TokenKeyEncryptionCreatorID
type TokenKeyEncryptionCreatorIDList []TokenKeyEncryptionCreatorID

// Len returns the number of IDs in the slice
func (tl TokenKeyEncryptionCreatorIDList) Len() int {
	return len(tl)
}

// Less returns true if the ID at i is less than at j
func (tl TokenKeyEncryptionCreatorIDList) Less(i, j int) bool {
	return tl[i] < tl[j]
}

// Swap will switch the IDs at i and j
func (tl TokenKeyEncryptionCreatorIDList) Swap(i, j int) {
	x := tl[i]
	tl[i] = tl[j]
	tl[j] = x
}

var errInvalidTokenKeyEncryptionCreator = errors.New("TokenKeyEncryptionCreator must not be nil")
var errInvalidTokenKeyEncryptionID = errors.New("TokenKeyEncryptionCreator must have a valid ID")
var errInvalidCreatorNoEncryptor = errors.New("TokenKeyEncryptionCreator must provide a non-nil Encryptor")
var errUnknownTokenKeyEncryptionID = errors.New("unknown TokenKeyEncryptionCreatorID")
var errAlgoTypeMismatch = errors.New("TokenKeyEncryptionCreator algo type is mismatched to the EncryptedObject algo type")

func init() {
	DefaultTokenKeyEncryptionFactory, _ = NewTokenKeyEncryptionFactory([]TokenKeyEncryptionCreator{
		&defaultGCMTokenKeyEncryption{},
	})
}

// DefaultTokenKeyEncryptionFactory is a TokenKeyEncryptionFactory pre-filled with
// with default TokenKeyEncryptionCreators
var DefaultTokenKeyEncryptionFactory TokenKeyEncryptionFactory

// NewTokenKeyEncryptionFactory returns an instance of TokenKeyEncryptionFactory,
// pre-filled with the specified set of TokenKeyEncryptionCreators
func NewTokenKeyEncryptionFactory(as []TokenKeyEncryptionCreator) (TokenKeyEncryptionFactory, error) {
	f := &tkeFactory{
		m: make(map[TokenKeyEncryptionCreatorID]TokenKeyEncryptionCreator),
	}

	for _, c := range as {
		if err := f.AddTokenKeyEncryptionCreator(c); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// tkeFactory is the implementation of the TokenKeyEncryptionFactory
// used by NewTokenKeyEncryptionFactory, simply providing a mapping
// to a TokenKeyEncryptionCreator via its ID
type tkeFactory struct {
	m map[TokenKeyEncryptionCreatorID]TokenKeyEncryptionCreator
	l sync.Mutex
}

// AddTokenKeyEncryptionCreator inserts the specified TokenKeyEncryptionCreator
// into the factory, replacing any existing creator with the same ID
func (f *tkeFactory) AddTokenKeyEncryptionCreator(c TokenKeyEncryptionCreator) error {
	f.l.Lock()
	defer f.l.Unlock()

	if c == nil {
		return errInvalidTokenKeyEncryptionCreator
	}

	_, err := NewAlgo(c.GetEncryptionAlgoType())
	if err != nil {
		return errUnknownAlgoType
	}

	if c.GetID() == "" {
		return errInvalidTokenKeyEncryptionID
	}

	_, err = c.GetTokenKeyEncryptor()
	if err != nil {
		return errInvalidCreatorNoEncryptor
	}

	// Will replace existing for the same ID
	f.m[c.GetID()] = c

	return nil
}

// GetTokenKeyEncryptionCreatorIDs returns a slice of the known IDs
func (f *tkeFactory) GetTokenKeyEncryptionCreatorIDs() TokenKeyEncryptionCreatorIDList {
	f.l.Lock()
	defer f.l.Unlock()

	ids := make(TokenKeyEncryptionCreatorIDList, 0, len(f.m))
	for k, _ := range f.m {
		ids = append(ids, k)
	}

	sort.Sort(ids)

	return ids
}

// GetTokenKeyDecryptor returns the Decryptor provided by the
// Creator with the specified ID.  The Decryptor may fail to initialise
// if the other arguments are not consistent with the creator
func (f *tkeFactory) GetTokenKeyDecryptor(i TokenKeyEncryptionCreatorID, key []byte, keys *EncryptedObject, factory AlgorithmFactory) (TokenKeyDecryptor, error) {
	f.l.Lock()
	defer f.l.Unlock()

	c, ok := f.m[i]
	if !ok {
		return nil, errUnknownTokenKeyEncryptionID
	}

	return c.GetTokenKeyDecryptor(key, keys, factory)
}

// GetTokenKeyEncryptor returns the Encryptor provided by the
// Creator with the specified ID.
func (f *tkeFactory) GetTokenKeyEncryptor(i TokenKeyEncryptionCreatorID) (TokenKeyEncryptor, error) {
	f.l.Lock()
	defer f.l.Unlock()

	c, ok := f.m[i]
	if !ok {
		return nil, errUnknownTokenKeyEncryptionID
	}

	return c.GetTokenKeyEncryptor()
}
