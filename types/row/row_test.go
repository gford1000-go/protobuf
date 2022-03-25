package row

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/gford1000-go/protobuf/types/cell"
	"github.com/gford1000-go/protobuf/types/encryption"
	"github.com/gford1000-go/protobuf/types/hashing"
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
	af        encryption.AlgorithmFactory
	d         encryption.TokenKeyDecryptor
	dtkf      encryption.TokenKeyEncryptionFactory
	envAlgo   encryption.AlgoType
	h         hashing.Hasher
	i         encryption.TokenKeyEncryptionCreatorID
	nToi      map[AttributeName]AttributeIdentifier
	iTon      map[AttributeIdentifier]AttributeName
	rowTokens [][]byte
	keyTokens []*keyToken
}

// init creates some dummy key tokens
func (e *controller) init(nRows int, attributeNames []AttributeName) {
	rand.Seed(99) // Used to consistently randomise whether a cell should be encrypted

	// Details for envelope encryption of the token->key map
	e.af = encryption.DefaultAlgoFactory
	e.envAlgo = encryption.GCM

	// Details for key token encryption and decryption
	e.dtkf = encryption.DefaultTokenKeyEncryptionFactory
	e.i = encryption.TokenKeyEncryptionCreatorID("DefaultGCM")

	// Attribute lookup map
	e.nToi = make(map[AttributeName]AttributeIdentifier)
	e.iTon = make(map[AttributeIdentifier]AttributeName)

	// The set of attributes in each row
	for i, n := range attributeNames {
		e.nToi[n] = AttributeIdentifier(i)
		e.iTon[AttributeIdentifier(i)] = n
	}

	// Row tokens - normally these would differentiate between
	// different access to the rows, here just have unique values
	e.rowTokens = make([][]byte, 0, nRows)
	for i := 0; i < nRows; i++ {
		e.rowTokens = append(e.rowTokens, []byte(fmt.Sprintf("%v", i)))
	}

	// Generate some psuedo random data in the rows
	e.keyTokens = make([]*keyToken, 0, nRows*len(attributeNames))
	for i := 0; i < nRows*len(attributeNames); i++ {
		kt := &keyToken{
			token:   []byte(fmt.Sprintf("%v", i)),
			encrypt: rand.Intn(2) == 1,
		}
		e.keyTokens = append(e.keyTokens, kt)
	}

	// Could have different hashing for each attribute but here
	// will just use the same hasher for all attributes
	e.h, _ = hashing.DefaultFactory.GetHasher(hashing.SHA256)
}

func (e *controller) getEnvelopeAlgorithm() (encryption.Algorithm, error) {
	return e.af.GetAlgorithm(encryption.GCM)
}

func (e *controller) createEnvelopeMasterKey() []byte {
	// Use the initialised algo and factory to generate
	// a new master key
	enve, _ := e.getEnvelopeAlgorithm()
	masterKey, _ := enve.CreateKey()
	return masterKey
}

func (e *controller) createEncryptor() (encryption.TokenKeyEncryptor, error) {
	// Use the initialised factory and algo to
	// create an encryptor
	return e.dtkf.GetTokenKeyEncryptor(e.i)
}

func (e *controller) createDecryptor(keys []byte, eo *encryption.EncryptedObject) {
	// Use the initialised factory and algo, the
	// encrypted token->key map, and the envelope master key,
	// to initialise an instance of the correct decryptor
	d, _ := e.dtkf.GetTokenKeyDecryptor(e.i, keys, eo, e.af)
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
	// Use the same decryptor for both attributes and rows.
	// Not recommended for production use
	return e.d
}

func (e *controller) GetHasher(name AttributeName) (hashing.Hasher, error) {
	// Single hasher used for all attributes.
	// Not recommended for production use
	return e.h, nil
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

	// Create the encryptor to be used on each cell
	e, _ := c.createEncryptor()

	// RowBuilder performs the marshaling
	rb, _ := NewRowBuilder(c, e, c)

	// Marshal the contents of the Row.
	// The row is serialised and encrypted, as well as
	// potentially encrypting specific cell values
	data, _, _ := rb.Marshal(rows[0].GetID(), rows[0].GetAll(), c)

	// Secure retrieval of token->key map
	envelopeMasterKey := c.createEnvelopeMasterKey()
	envelopeAlgo, _ := c.getEnvelopeAlgorithm()
	keys, _ := e.GetKeys(envelopeMasterKey, envelopeAlgo)

	// Transfer of token->key map to a decryptor.
	// Access control is normally applied at this point, which
	// limits the visibility of data to the decryptor
	c.createDecryptor(envelopeMasterKey, keys)

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
