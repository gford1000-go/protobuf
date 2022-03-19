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

type defaultGCMTokenKeyEncryption struct {
}

func (d *defaultGCMTokenKeyEncryption) GetID() TokenKeyEncryptionCreatorID {
	return "DefaultGCM"
}

func (d *defaultGCMTokenKeyEncryption) GetEncryptionAlgoType() AlgoType {
	return GCM
}

func (d *defaultGCMTokenKeyEncryption) GetTokenKeyDecryptor(key []byte, keys *EncryptedObject, factory AlgorithmFactory) (TokenKeyDecryptor, error) {
	if keys == nil {
		return nil, errObjectIsNil
	}

	if factory == nil {
		return nil, errFactoryIsNil
	}

	a, err := ParseAlgo(keys.A)
	if err != nil {
		return nil, err
	}

	if a != d.GetEncryptionAlgoType() {
		return nil, errAlgoTypeMismatch
	}

	algo, err := factory.GetAlgorithm(a)
	if err != nil {
		return nil, err
	}

	if algo.GetDecryptor() == nil {
		return nil, errNoDecryptor
	}

	b, err := algo.GetDecryptor().Decrypt(key, keys.V)
	if err != nil {
		return nil, err
	}

	k := &Keys{}
	err = proto.Unmarshal(b, k)
	if err != nil {
		return nil, err
	}

	return &defaultTokenKeyDecryptor{
		a:    algo,
		keys: k.GetKeys(),
	}, nil
}

func (d *defaultGCMTokenKeyEncryption) GetTokenKeyEncryptor() (TokenKeyEncryptor, error) {
	a, err := DefaultAlgoFactory.GetAlgorithm(d.GetEncryptionAlgoType())
	if err != nil {
		return nil, errUnknownAlgoType
	}
	return &defaultTokenKeyEncryptor{
		a:    a,
		keys: make(map[string][]byte),
	}, nil
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
