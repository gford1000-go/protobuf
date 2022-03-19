package encryption

import (
	"bytes"
	"testing"
)

func TestEncryption(t *testing.T) {

	data := [][]byte{
		{},
		[]byte("Hello World"),
		{00, 11, 22, 33},
	}

	k, _ := newAESKey()

	gcmE, _ := newGCMEncryptor(k)
	gcmD, _ := newGCMEncryptor(k)

	for _, d := range data {

		e, err := gcmE.Encrypt(d)
		if err != nil {
			t.Errorf("unexpected error - %v", err)
		}

		p, err := gcmD.Decrypt(e)
		if err != nil {
			t.Errorf("unexpected error - %v", err)
		}

		if !bytes.Equal(d, p) {
			t.Errorf("unexpected mismatch")
		}
	}
}

func TestNonEqualEncryption(t *testing.T) {

	testData := [][]byte{
		{},
		[]byte("Hello World"),
	}

	k, _ := newAESKey()

	gcmE, _ := newGCMEncryptor(k)

	for _, data := range testData {

		e, _ := gcmE.Encrypt(data)

		for i := 0; i < 100; i++ {
			ei, _ := gcmE.Encrypt(data)

			// Should not match
			if bytes.Equal(e, ei) {
				t.Errorf("unexpected match")
			}
		}
	}
}

func TestNonEqualPlaintext(t *testing.T) {

	testData := [][]byte{
		{},
		[]byte("Hello World"),
	}

	k, _ := newAESKey()

	gcmE, _ := newGCMEncryptor(k)

	for _, p := range testData {

		e, _ := gcmE.Encrypt(p)

		// Should not match
		if bytes.Equal(e, p) {
			t.Errorf("unexpected match")
		}
	}
}
