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

// dummyKeyManager holds the tokens used in the processing
type dummyKeyManager struct {
	keyTokens []*keyToken
}

// init creates the specified set of tokens
func (d *dummyKeyManager) init(n int) {
	d.keyTokens = make([]*keyToken, 0, n)
	for i := 0; i < n; i++ {
		kt := &keyToken{
			token:   []byte(fmt.Sprintf("%v", rand.Intn(n*1000))),
			encrypt: rand.Intn(2) == 1,
		}
		d.keyTokens = append(d.keyTokens, kt)
	}
}

func (d *dummyKeyManager) GetToken() []byte {
	// Would do something clever here
	// for now, return random selection
	return d.keyTokens[rand.Intn(len(d.keyTokens))].token
}

func (d *dummyKeyManager) Encrypt(keyToken []byte) bool {
	for _, k := range d.keyTokens {
		if bytes.Equal(k.token, keyToken) {
			return k.encrypt
		}
	}
	return false
}

func ExampleCellBuilder() {
	rand.Seed(time.Now().Unix())

	// This emulates logic that has assigned keyTokens to the contents of each Cell
	km := &dummyKeyManager{}
	km.init(100)

	// Create the Hasher that returns the hash of the cell's value
	h, _ := hashing.DefaultFactory.GetHasher(hashing.SHA256)

	e := encryption.NewGCMTokenKeyEncryptor()

	cb, _ := NewCellBuilder(e, h)

	// This is the cell's value
	i := createValue()

	// Marshal the value, using the determiners to assign the
	// correct keyToken and apply encryption as required - this
	// would typically need awareness of other data but here we
	// use the dummyKeyManager to assign randomly
	data, _, _ := cb.Marshal(i, km, km)

	// Get the keys used to encrypt the cell
	// need to supply a master key with which they are secured
	masterKey := []byte("0123456789abcdef") // Should be random
	k, _ := e.GetKeys(masterKey)

	// The decryptor is initialised by providing the secured keys
	// and the masterKey
	d, _ := encryption.NewGCMTokenKeyDecryptor(masterKey, k)

	// Create a cell parser, supplying the decryptor
	cp, _ := NewCellParser(d)

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
