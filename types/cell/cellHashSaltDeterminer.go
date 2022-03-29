package cell

type CellHashSaltDeterminer interface {
	GetSalt() []byte
}
