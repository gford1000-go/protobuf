package encryption

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

var errObjectIsNil = errors.New("EncryptedObject must not be nil")
var errFactoryIsNil = errors.New("AlgorithmFactory must not be nil")
var errAlgoMismatch = errors.New("AlgoType mismatch")
var errMissingKeyToken = errors.New("keyToken not found")
var errNoDecryptor = errors.New("Algorithm returns nil Decryptor")
var errInvalidKeyToken = errors.New("keyToken not valid")
var errEncryptionError = errors.New("error during encryption")
var errInvalidDecryptorCreator = errors.New("TokenKeyDecryptorCreator must not be nil")
var errInvalidEncryptorCreator = errors.New("TokenKeyEncryptorCreator must not be nil")
var errInvalidAlgorithmFactory = errors.New("AlgorithmFactory must not be nil")

// TokenKeyDecryptorCreator returns an initialised TokenKeyDecryptor
type TokenKeyDecryptorCreator func(a Algorithm, keys map[string][]byte) TokenKeyDecryptor

// TokenKeyEncryptorCreator returns an initialised TokenKeyEncryptor
type TokenKeyEncryptorCreator func(a Algorithm) TokenKeyEncryptor

// NewTokenKeyEncryptionCreator provides a construction mechanism
// to create instances of TokenKeyEncryptionCreator
func NewTokenKeyEncryptionCreator(
	id TokenKeyEncryptionCreatorID,
	a AlgoType,
	d TokenKeyDecryptorCreator,
	e TokenKeyEncryptorCreator,
	f AlgorithmFactory) (TokenKeyEncryptionCreator, error) {

	if id == "" {
		return nil, errInvalidTokenKeyEncryptionID
	}
	if d == nil {
		return nil, errInvalidDecryptorCreator
	}
	if e == nil {
		return nil, errInvalidEncryptorCreator
	}
	if f == nil {
		return nil, errInvalidAlgorithmFactory
	}

	return &defaultTokenKeyEncryption{
		i: id,
		a: a,
		d: d,
		e: e,
		f: f,
	}, nil
}

type defaultTokenKeyEncryption struct {
	i TokenKeyEncryptionCreatorID
	a AlgoType
	d TokenKeyDecryptorCreator
	e TokenKeyEncryptorCreator
	f AlgorithmFactory
}

func (d *defaultTokenKeyEncryption) GetID() TokenKeyEncryptionCreatorID {
	return d.i
}

func (d *defaultTokenKeyEncryption) GetEncryptionAlgoType() AlgoType {
	return d.a
}

func (d *defaultTokenKeyEncryption) GetTokenKeyDecryptor(key []byte, keys *EncryptedObject, factory AlgorithmFactory) (TokenKeyDecryptor, error) {
	if keys == nil {
		return nil, errObjectIsNil
	}

	if factory == nil {
		return nil, errFactoryIsNil
	}

	// Determine the AlgoType with which the EncryptedObject
	// has been encrypted
	a, err := ParseAlgo(keys.A)
	if err != nil {
		return nil, err
	}

	// Retrieve the Algorithm from the specified factory
	algo, err := factory.GetAlgorithm(a)
	if err != nil {
		return nil, err
	}

	if algo.GetDecryptor() == nil {
		return nil, errNoDecryptor
	}

	// Decrypt the keys using the supplied key, and the
	// Algorithm from the supplied factory
	b, err := algo.GetDecryptor().Decrypt(key, keys.V)
	if err != nil {
		return nil, err
	}

	// Should have decrypted Keys object, so
	// attempt deserialisation
	k := &Keys{}
	err = proto.Unmarshal(b, k)
	if err != nil {
		return nil, err
	}

	// Now can retrieve the Algorithm for this Creator,
	// from the factory specified during Creator initialisation
	algo, err = d.f.GetAlgorithm(d.GetEncryptionAlgoType())
	if err != nil {
		return nil, errUnknownAlgoType
	}

	// Return the TokenKeyDecryptor, prefilled with keys
	// and Algorithm to be used
	return d.d(algo, k.GetKeys()), nil
}

func (d *defaultTokenKeyEncryption) GetTokenKeyEncryptor() (TokenKeyEncryptor, error) {
	// Retrieve the Algorithm for this Creator
	algo, err := d.f.GetAlgorithm(d.GetEncryptionAlgoType())
	if err != nil {
		return nil, errUnknownAlgoType
	}

	// Return the TokenKeyEncryptor, prefilled with the
	// Algorithm it should use
	return d.e(algo), nil
}

// newDefaultGCMCreator creates the default implementation of a
// Creator for the GCM encryption, which simply errors if the
// token is missing or decryption fails
func newDefaultGCMCreator() (TokenKeyEncryptionCreator, error) {
	d := func(a Algorithm, keys map[string][]byte) TokenKeyDecryptor {
		return &defaultTokenKeyDecryptor{
			a:    a,
			keys: keys,
		}
	}

	e := func(a Algorithm) TokenKeyEncryptor {
		return &defaultTokenKeyEncryptor{
			a:    a,
			keys: make(map[string][]byte),
		}
	}

	return NewTokenKeyEncryptionCreator(
		TokenKeyEncryptionCreatorID("DefaultGCM"),
		GCM,
		d,
		e,
		DefaultAlgoFactory,
	)
}

// defaultTokenKeyDecryptor implements TokenKeyDecryptor for decryption
type defaultTokenKeyDecryptor struct {
	keys map[string][]byte
	a    Algorithm
}

// DecryptFromToken attempts to decrypt using the key associated with the token.
func (g *defaultTokenKeyDecryptor) DecryptFromToken(keyToken []byte, algo AlgoType, ciphertext []byte) ([]byte, error) {
	if algo != g.a.GetType() {
		return nil, errAlgoMismatch
	}

	k, ok := g.keys[string(keyToken)]
	if !ok {
		return nil, errMissingKeyToken
	}

	return g.a.GetDecryptor().Decrypt(k, ciphertext)
}

// defaultTokenKeyEncryptor implements TokenKeyEncryptor for encryption
type defaultTokenKeyEncryptor struct {
	a    Algorithm
	keys map[string][]byte
}

// GetKeys returns the map of key tokens to their associated keys,
// encrypted using the specified key and the specified Algorithm
func (g *defaultTokenKeyEncryptor) GetKeys(key []byte, a Algorithm) (*EncryptedObject, error) {

	k := &Keys{
		Keys: g.keys,
	}

	return NewEncryptedObject(key, k, a.GetEncryptor())
}

// EncryptFromToken attempts to encrypt using the key associated with the token.
func (g *defaultTokenKeyEncryptor) EncryptFromToken(keyToken []byte, plaintext []byte) ([]byte, AlgoType, error) {
	if len(keyToken) == 0 {
		return nil, Unknown, errInvalidKeyToken
	}

	k, ok := g.keys[string(keyToken)]
	if !ok {
		var err error
		k, err = g.a.CreateKey()
		if err != nil {
			return nil, Unknown, err
		}
		g.keys[string(keyToken)] = k
	}

	return g.a.GetEncryptor().Encrypt(k, plaintext)
}
