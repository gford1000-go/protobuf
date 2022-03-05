package cell

// TokenKeyRetriever returns the appropriate keyToken
type TokenKeyRetriever interface {
	GetToken() []byte
}
