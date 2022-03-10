package row

import (
	"github.com/gford1000-go/protobuf/types/encryption"
)

// RowDecryptionInterfacesRetriever returns the appropriate interfaces
// to be able to decrypt a row of attribute data
type RowDecryptionInterfacesRetriever interface {
	GetRowTokenKeyDecryptor() encryption.TokenKeyDecryptor
	GetCellsTokenKeyDecryptor(rowID RowID) encryption.TokenKeyDecryptor
}
