package row

import "github.com/gford1000-go/protobuf/types/cell"

// RowEncryptionInterfacesRetriever returns the appropriate interfaces
// to be able to serialise a row of attribute data
type RowEncryptionInterfacesRetriever interface {
	GetRowToken() ([]byte, error)
	GetValueEncryptionDeterminer() cell.ValueEncryptionDeterminer
	GetTokenKeyRetriever(attribute AttributeName) cell.TokenKeyRetriever
}
