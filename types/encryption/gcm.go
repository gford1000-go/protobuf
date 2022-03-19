package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// newAESKey returns a new key for AES encryption
func newAESKey() ([]byte, error) {
	k := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil, fmt.Errorf("error creating key - %v", err)
	}
	return k, nil
}

// newGCMEncryptor creates a new instance, initialised with the key
func newGCMEncryptor(key []byte) (*gcmEncryptor, error) {

	if len(key) != aes.BlockSize {
		return nil, fmt.Errorf("key is wrong size - should be %v bytes", aes.BlockSize)
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	return &gcmEncryptor{
		c: c,
	}, nil
}

// gcmEncryptor provides GCM encryption and decryption
type gcmEncryptor struct {
	c cipher.Block
}

// Encrypt returns the ciphertext for the plaintext
func (g *gcmEncryptor) Encrypt(plaintext []byte) ([]byte, error) {

	gcm, err := cipher.NewGCM(g.c)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM - %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("error creating Nonce - %v", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

// Decrypt returns the plaintext for the ciphertext
func (g *gcmEncryptor) Decrypt(ciphertext []byte) ([]byte, error) {

	gcm, err := cipher.NewGCM(g.c)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM - %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("data is wrong size - should be greater than %v bytes", gcm.NonceSize())
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting - %v", err)
	}

	return plaintext, nil
}
