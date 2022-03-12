package hashing

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDefault(t *testing.T) {

	ht := SHA256
	salt := []byte("12345")

	h, err := DefaultFactory.GetHasherWithSalt(ht, salt)
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	if h.GetHashType() != ht {
		t.Fatal("Mismatch in HashType")
	}

	if !bytes.Equal(h.GetSalt(), salt) {
		t.Fatal("Mismatch in salts")
	}

	type testData struct {
		i interface{}
		h string
	}

	data := []testData{
		{int64(1), "dd712114fb283417de4da3512e17486adbda004060d0d1646508c8a2740d29b4"},
		{float32(-9992.3), "3568d479b826f1750d31755023a387a623cbbf97358841fbccf2af2e8f833880"},
	}

	for _, d := range data {
		if fmt.Sprintf("%x", h.Hash(d.i).H) != d.h {
			t.Fatalf("Mismatch in hash for %v\n", d.i)
		}
	}
}

func TestRepeated(t *testing.T) {

	salt := []byte("12345")

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, salt)
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h12345 := "dd712114fb283417de4da3512e17486adbda004060d0d1646508c8a2740d29b4"
	h1 := fmt.Sprintf("%x", h.Hash(int64(1)).H)
	h2 := fmt.Sprintf("%x", h.Hash(int64(1)).H)

	if h1 != h2 || h1 != h12345 || h2 != h12345 {
		t.Fatal("Mismatch in repeated hashing of the same value")
	}
}

func TestNotSameHash(t *testing.T) {

	hasher1, err := DefaultFactory.GetHasher(SHA256)
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	hasher2, err := DefaultFactory.GetHasher(SHA256)
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h12345 := "dd712114fb283417de4da3512e17486adbda004060d0d1646508c8a2740d29b4"
	h1 := fmt.Sprintf("%x", hasher1.Hash(int64(1)).H)
	h2 := fmt.Sprintf("%x", hasher2.Hash(int64(1)).H)

	if h1 == h2 || h1 == h12345 || h2 == h12345 {
		fmt.Println(h12345)
		fmt.Println(h1)
		fmt.Println(h2)
		t.Fatal("Different hashers give same hash for the same value")
	}
}

func TestUnknownHashType(t *testing.T) {
	hasher, err := DefaultFactory.GetHasher(HashType(99))
	if hasher != nil {
		t.Fatal("Invalid Hasher returned")
	}
	if err != errUnknownHashType {
		t.Fatal("Incorrect error returned")
	}
}

func TestNilSalt(t *testing.T) {

	ht := SHA256

	h, err := DefaultFactory.GetHasherWithSalt(ht, nil)
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	if h.GetHashType() != ht {
		t.Fatal("Mismatch in HashType")
	}

	if !bytes.Equal(h.GetSalt(), nil) {
		t.Fatal("Mismatch in salts")
	}

	type testData struct {
		i interface{}
		h string
	}

	data := []testData{
		{int64(1), "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"},
		{float32(-9992.3), "5761622cb1454cd71c280cf82307861257cc2784d2505421821dd05a926ef1e7"},
	}

	for _, d := range data {
		if fmt.Sprintf("%x", h.Hash(d.i).H) != d.h {
			t.Fatalf("Mismatch in hash for %v\n", d.i)
		}
	}
}

func TestNilHasherCreator(t *testing.T) {
	err := DefaultFactory.AddHasherCreator(nil)
	if err != errInvalidHasherCreator {
		t.Fatal("Incorrect error returned")
	}
}

type testHC struct {
}

func (thc *testHC) New(salt []byte) Hasher {
	return nil
}

func TestBadHasherCreator(t *testing.T) {
	_, err := NewFactory([]HasherCreator{&testHC{}})
	if err != errCreatorDoesNotEmitHasher {
		t.Fatal("Incorrect error returned")
	}
}
