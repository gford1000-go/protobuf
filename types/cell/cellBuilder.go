package cell

import (
	"errors"

	"github.com/gford1000-go/protobuf/types/encryptable_value"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/hashing"
	"github.com/gford1000-go/protobuf/types/value"
	"google.golang.org/protobuf/proto"
)

var errMissingEncryptor = errors.New("TokenKeyEncryptor must not be nil")
var errMissingHasherFactory = errors.New("HasherFactory must not be nil")
var errHasherFactoryReturnsNilHasher = errors.New("HasherFactory returned a nil Hasher")
var errMissingSaltDeterminer = errors.New("CellHashSaltDeterminer must not be nil")
var errMissingTokenKeyDeterminer = errors.New("TokenKeyDeterminer must not be nil")
var errMissingValueEncryptionDeterminer = errors.New("ValueEncryptionDeterminer must not be nil")

// NewCellBuilder returns an initialised instance of CellBuilder
func NewCellBuilder(encryptor encryption.TokenKeyEncryptor, factory hashing.HasherFactory, hashtype hashing.HashType) (*CellBuilder, error) {
	if encryptor == nil {
		return nil, errMissingEncryptor
	}
	if factory == nil {
		return nil, errMissingHasherFactory
	}
	h, err := factory.GetHasher(hashtype)
	if err != nil {
		return nil, err
	}
	if h == nil {
		return nil, errHasherFactoryReturnsNilHasher
	}
	return &CellBuilder{e: encryptor, f: factory, ht: hashtype}, nil
}

// CellBuilder serialises Cells, optionally applying encryption
type CellBuilder struct {
	e  encryption.TokenKeyEncryptor
	f  hashing.HasherFactory
	ht hashing.HashType
}

// Marshal serialises the specified value, dependent on the value's token and whether the token signifies the value should be encrypted or not.
// If a CellTokenDeterminer is provided, then
// Marshal also returns a Hash instance representing the value, generated using the salt specified by CellHashSaltDeterminer
func (cb *CellBuilder) Marshal(i interface{}, s CellHashSaltDeterminer, td CellTokenDeterminer, d TokenKeyRetriever, ve ValueEncryptionDeterminer) ([]byte, *hashing.Hash, error) {
	if s == nil {
		return nil, nil, errMissingSaltDeterminer
	}

	// Build a hasher with the appropriate salt for the Cell, and generate the hash
	hasher, err := cb.f.GetHasherWithSalt(cb.ht, s.GetSalt())
	if err != nil {
		return nil, nil, err
	}
	h := hasher.Hash(i)

	var c *Cell
	if td != nil {
		token, ok := td.GetToken(h)

		if !ok {
			// Hash not seen before - create the Cell and then obtain a token for it
			c, err = ToCell(i, d, ve, cb.e)
			if err != nil {
				return nil, nil, err
			}

			token, err = td.CreateToken(c, h)
			if err != nil {
				return nil, nil, err
			}
		}

		// Returned cell only contains the token
		c = &Cell{C: &Cell_Token{Token: token}}
	} else {
		// Cell contains an EncryptableValue
		c, err = ToCell(i, d, ve, cb.e)
		if err != nil {
			return nil, nil, err
		}
	}

	data, err := proto.Marshal(c)
	if err != nil {
		return nil, nil, err
	}

	return data, h, nil
}

// ToCell returns a Cell populated with an EncryptableValue, with a key token provided the by TokenKeyRetriever,
// and encryption applied by the TokenKeyEncryptor if indicated by the ValueEncryptionDeterminer
func ToCell(i interface{}, d TokenKeyRetriever, ve ValueEncryptionDeterminer, e encryption.TokenKeyEncryptor) (*Cell, error) {
	if d == nil {
		return nil, errMissingTokenKeyDeterminer
	}
	if e == nil {
		return nil, errMissingEncryptor
	}
	if ve == nil {
		return nil, errMissingValueEncryptionDeterminer
	}

	v, err := value.NewValue(i)
	if err != nil {
		return nil, err
	}

	var keyToken []byte = d.GetToken()
	var encrypt bool = ve.Encrypt(keyToken)

	ev, err := encryptable_value.NewEncryptableValue(keyToken, v, encrypt, e)
	if err != nil {
		return nil, err
	}

	c := &Cell{C: &Cell_Scalar{Scalar: ev}}

	return c, nil
}
