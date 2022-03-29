package cell

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/hashing"
	"github.com/gford1000-go/protobuf/types/value"
)

func createValue() interface{} {
	data := []interface{}{
		"Hello World",
		int32(1234),
		float64(9999.99),
		true,
		[]interface{}{
			int64(5678),
			float32(-12.34),
			"Lunchtime",
			false,
			nil,
			[]interface{}{},
		},
		map[string]interface{}{
			"a": "Dinnertime",
			"b": float64(-7e10),
			"c": nil,
			"d": map[string]interface{}{
				"e": []byte("Breakfast"),
			},
			"e": []interface{}{
				false, true,
			},
		},
	}

	return data[rand.Intn(len(data))]
}

// keyToken holds a token and whether encryption is required
type keyToken struct {
	token   []byte
	encrypt bool
}

// packer holds the tokens used in the processing
type packer struct {
	keyTokens []*keyToken
}

// init creates the specified set of tokens
func (p *packer) init(n int) {
	p.keyTokens = make([]*keyToken, 0, n)
	for i := 0; i < n; i++ {
		kt := &keyToken{
			token:   []byte(fmt.Sprintf("%v", rand.Intn(n*1000))),
			encrypt: rand.Intn(2) == 1,
		}
		p.keyTokens = append(p.keyTokens, kt)
	}
}

func (p *packer) GetSalt() []byte {
	// Provide a random salt for production use
	return []byte("xyz")
}

func (p *packer) GetToken() []byte {
	// Would do something clever here
	// for now, return random selection
	return p.keyTokens[rand.Intn(len(p.keyTokens))].token
}

func (p *packer) Encrypt(keyToken []byte) bool {
	for _, k := range p.keyTokens {
		if bytes.Equal(k.token, keyToken) {
			return k.encrypt
		}
	}
	return false
}

func ExampleCellBuilder() {
	rand.Seed(time.Now().Unix())

	// This emulates logic that has assigned keyTokens to the contents of each Cell
	p := &packer{}
	p.init(100)

	// Retrieve the TokenKeyEncryptor for the specified ID
	id := encryption.TokenKeyEncryptionCreatorID("DefaultGCM")
	e, _ := encryption.DefaultTokenKeyEncryptionFactory.GetTokenKeyEncryptor(id)

	// Create a CellBuilder which will use SHA256 from the default factory to create hashes
	cb, _ := NewCellBuilder(e, hashing.DefaultFactory, hashing.SHA256)

	// This is the cell's value
	i := createValue()

	// Marshal the value, using the determiners to assign the
	// correct keyToken and apply encryption as required - this
	// would typically need awareness of other data but here we
	// use the dummyKeyManager to assign randomly
	data, _, _ := cb.Marshal(i, p, nil, p, p)

	// Envelope key and algorithm, for keys transfer
	gcm, _ := encryption.DefaultAlgoFactory.GetAlgorithm(encryption.GCM)
	masterKey, _ := gcm.CreateKey()

	// Get the keys used to encrypt the cell
	// need to supply a master key with which they are secured
	k, _ := e.GetKeys(masterKey, gcm)

	// Retrieve the TokenKeyDecryptor for the specified ID
	d, _ := encryption.DefaultTokenKeyEncryptionFactory.GetTokenKeyDecryptor(id, masterKey, k, encryption.DefaultAlgoFactory)

	// Create a cell parser, supplying the decryptor
	cp, _ := NewCellParser(d, nil)

	// Parse a Cell back to its constituent Value(s).  For now the
	// Cell only supports scalar entries (i.e. only one Value, which
	// itself might be a complex type), but will in future support
	// lists and maps of Values as well
	v, _ := cp.Parse(data)

	// Parse the Value to the underlying go type
	ii, _ := value.ParseValue(v)

	fmt.Println(fmt.Sprint(i) == fmt.Sprint(ii))
	// Output: true
}
