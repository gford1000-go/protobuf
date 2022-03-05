package cell

// ValueEncryptionDeterminer returns whether encryption should
// be applied for a given keyToken
type ValueEncryptionDeterminer interface {
	Encrypt(keyToken []byte) bool
}
