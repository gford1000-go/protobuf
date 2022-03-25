package row

import "github.com/gford1000-go/protobuf/types/hashing"

// AttributeHashRetriever returns the Hasher associated with the specified
// AttributeName, or an error if no Hasher is defined
type AttributeHashRetriever interface {
	GetHasher(name AttributeName) (hashing.Hasher, error)
}
