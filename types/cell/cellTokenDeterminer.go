package cell

import "github.com/gford1000-go/protobuf/types/hashing"

type CellTokenDeterminer interface {
	GetToken(h *hashing.Hash) (uint64, bool)
	CreateToken(c *Cell, h *hashing.Hash) (uint64, error)
}
