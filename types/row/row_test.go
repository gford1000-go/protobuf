package row

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/gford1000-go/protobuf/types/cell"
	"github.com/gford1000-go/protobuf/types/encryption"
)

func createRows() []Row {
	// Return just a single row for the example
	data := []Row{
		{
			rowId: 0,
			atts: map[AttributeName]interface{}{
				"a": "Hello World",
				"b": int32(1234),
				"c": float64(9999.99),
				"d": true,
				"e": []interface{}{
					int64(5678),
					float32(-12.34),
					"Lunchtime",
					false,
					nil,
					[]interface{}{},
				},
				"f": map[string]interface{}{
					"aa": "Dinnertime",
					"bb": float64(-7e10),
					"cc": nil,
					"dd": map[string]interface{}{
						"e": []byte("Breakfast"),
					},
					"ee": []interface{}{
						false, true,
					},
				},
			},
		},
	}

	return data
}

// keyToken holds a token and whether encryption is required
type keyToken struct {
	token   []byte
	encrypt bool
}

// controller implements all interfaces needed to marshal and parse a row
type controller struct {
	d         encryption.TokenKeyDecryptor
	nToi      map[AttributeName]AttributeIdentifier
	iTon      map[AttributeIdentifier]AttributeName
	rowTokens [][]byte
	keyTokens []*keyToken
}

// init creates some dummy key tokens
func (e *controller) init(nRows int, attributeNames []AttributeName) {
	rand.Seed(99) // Used to consistently randomise whether a cell should be encrypted

	e.nToi = make(map[AttributeName]AttributeIdentifier)
	e.iTon = make(map[AttributeIdentifier]AttributeName)

	for i, n := range attributeNames {
		e.nToi[n] = AttributeIdentifier(i)
		e.iTon[AttributeIdentifier(i)] = n
	}

	e.rowTokens = make([][]byte, 0, nRows)
	for i := 0; i < nRows; i++ {
		e.rowTokens = append(e.rowTokens, []byte(fmt.Sprintf("%v", i)))
	}

	e.keyTokens = make([]*keyToken, 0, nRows*len(attributeNames))
	for i := 0; i < nRows*len(attributeNames); i++ {
		kt := &keyToken{
			token:   []byte(fmt.Sprintf("%v", i)),
			encrypt: rand.Intn(2) == 1,
		}
		e.keyTokens = append(e.keyTokens, kt)
	}
}

func (e *controller) createDecryptor(keys []byte, eo *encryption.EncryptedObject) {
	d, _ := encryption.NewGCMTokenKeyDecryptor(keys, eo)
	e.d = d
}

func (e *controller) FindName(i AttributeIdentifier) (AttributeName, error) {
	n, ok := e.iTon[i]
	if !ok {
		return AttributeName(""), fmt.Errorf("unknown identifier %v", i)
	}
	return n, nil
}

func (e *controller) FindIdentifier(n AttributeName) (AttributeIdentifier, error) {
	i, ok := e.nToi[n]
	if !ok {
		return AttributeIdentifier(0), fmt.Errorf("unknown name %v", n)
	}
	return i, nil
}

func (e *controller) GetRowToken() ([]byte, error) {
	return e.rowTokens[rand.Intn(len(e.rowTokens))], nil
}

func (e *controller) GetValueEncryptionDeterminer() cell.ValueEncryptionDeterminer {
	return e
}

func (e *controller) GetTokenKeyRetriever(attribute AttributeName) cell.TokenKeyRetriever {
	return e
}

func (e *controller) GetToken() []byte {
	return e.keyTokens[rand.Intn(len(e.keyTokens))].token
}

func (e *controller) Encrypt(keyToken []byte) bool {
	for _, k := range e.keyTokens {
		if bytes.Equal(k.token, keyToken) {
			return k.encrypt
		}
	}
	return false
}

func (e *controller) GetRowTokenKeyDecryptor() encryption.TokenKeyDecryptor {
	return e.d
}
func (e *controller) GetCellsTokenKeyDecryptor(rowID RowID) encryption.TokenKeyDecryptor {
	return e.d
}

func ExampleRowBuilder() {

	// This is an example set of rows, with each row
	// comprising a map of names -> values, where each value
	// has the same type for a given name
	rows := createRows()

	// This class implements all the interfaces required to serialise
	// and deserialise the data in the row
	c := &controller{}
	c.init(len(rows), rows[0].GetAttributeNames())

	// For this example, ensure encryptors and decryptors observe
	// the same map of token->key, so that decryption is successful.
	// In normal use, the data requestor will only have a subset of
	// the full map, which controls their access to data.
	e := encryption.NewGCMTokenKeyEncryptor()

	cb, _ := NewRowBuilder(e, c)

	// Marshal the contents of the Row.
	// The row is serialised and encrypted, as well as
	// potentially encrypting specific cell values
	data, _ := cb.Marshal(rows[0].GetID(), rows[0].GetAll(), c)

	// Transfer the encryption keys to the controller object
	masterKey := []byte("0123456789abcdef") // Should be random
	keys, _ := e.GetKeys(masterKey)
	c.createDecryptor(masterKey, keys)

	cp, _ := NewRowParser(c, c)

	// Parse the encrypted row bytes, unpacking to return a
	// Row instance, with access to the attributes and their values
	//
	// In this example, the controller ensures that all row tokens
	// and cell tokens are available, so the data can be fully parsed.
	// In normal use, the data requestor will only have a subset of the
	// row tokens (i.e. only allowed access to some rows), and where
	// access is available, may only have access to the values of
	// certain attributes of the row (and this may vary, row by row).
	row, _ := cp.Parse(data)

	// Retrieve the map of all attributes
	atts1 := row.GetAll()

	fmt.Println(fmt.Sprint(rows[0].GetAll()) == fmt.Sprint(atts1))
	// Output: true
}
