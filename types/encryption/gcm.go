package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// NewAESKey returns a new key for AES encryption
func NewAESKey() ([]byte, error) {
	k := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil, fmt.Errorf("error creating key - %v", err)
	}
	return k, nil
}

// NewGCMEncryptor creates a new instance, initialised with the key
func NewGCMEncryptor(key []byte) (*GCMEncryptor, error) {

	if len(key) != aes.BlockSize {
		return nil, fmt.Errorf("key is wrong size - should be %v bytes", aes.BlockSize)
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	return &GCMEncryptor{
		c: c,
	}, nil
}

// GCMEncryptor provides GCM encryption and decryption
type GCMEncryptor struct {
	c cipher.Block
}

// Encrypt returns the ciphertext for the plaintext
func (g *GCMEncryptor) Encrypt(plaintext []byte) ([]byte, error) {

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
func (g *GCMEncryptor) Decrypt(ciphertext []byte) ([]byte, error) {

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
