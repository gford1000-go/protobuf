package cell

// CellTokenRetriever returns the Cell for a given token, or
// returns an error if the token is unknown
type CellTokenRetriever interface {
	GetCell(token uint64) (*Cell, error)
}
