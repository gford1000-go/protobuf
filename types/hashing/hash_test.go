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

	type testStruct struct {
		a int64
		b string
	}

	data := []testData{
		{nil, "589afb9f637093785273a735592d92a92e13a52f07a128770cde82c0e45b656a"},
		{int64(1), "dd712114fb283417de4da3512e17486adbda004060d0d1646508c8a2740d29b4"},
		{float32(-9992.3), "3568d479b826f1750d31755023a387a623cbbf97358841fbccf2af2e8f833880"},
		{[]int64{1, 2, 3, 4, 5}, "cd5b761f3877aacab78715384f023546cbf462b1ece14de26f7e1ebad1ac591e"},
		{map[string]interface{}{"a": 1, "b": "Hello", "c": []float64{2.3, 4.5}}, "50c7d2d1b89d346e1386244aa4d011c963795036064c6367838277f3c421e20d"},
		{map[string]interface{}{"b": "Hello", "a": 1, "c": []float64{2.3, 4.5}}, "50c7d2d1b89d346e1386244aa4d011c963795036064c6367838277f3c421e20d"},
		{[]*testStruct{{a: 1, b: "Hi"}, {a: 2, b: "There"}}, "20dfd7869c5d4609cf7afdf09ae8dc0967e554e10b8d8200b30834405fd7e4ca"},
		{[]*testStruct{{a: 2, b: "There"}, {a: 1, b: "Hi"}}, "20dfd7869c5d4609cf7afdf09ae8dc0967e554e10b8d8200b30834405fd7e4ca"},
		{[]*[]int64{{7, 8, 9, 0}, {1, 2, 3, 4}}, "ee2d93a36e99ca976ff5ec139e8c5e595f92160c203e6215d74f75f29cebe413"},
		{map[int]*testStruct{1: {a: 1, b: "Hi"}, 2: {a: 2, b: "There"}}, "b0779c149603990ff73355a065894126966bdc129ffac6c4db07f5fb760b1e5a"},
		{map[int]*testStruct{2: {a: 2, b: "There"}, 1: {a: 1, b: "Hi"}}, "b0779c149603990ff73355a065894126966bdc129ffac6c4db07f5fb760b1e5a"},
	}

	for _, d := range data {
		if fmt.Sprintf("%x", h.Hash(d.i).H) != d.h {
			fmt.Printf("%x\n", h.Hash(d.i).H)
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

func TestSlicesEqual(t *testing.T) {

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, []byte("12345"))
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h1 := h.Hash([]int64{1, 2, 3}).H
	h2 := h.Hash([]int64{2, 3, 1}).H
	h3 := h.Hash([]int64{2, 1, 3}).H

	if !bytes.Equal(h1, h2) || !bytes.Equal(h1, h3) {
		t.Fatal("Slice equialence error")
	}
}

func TestSlicesDifferent(t *testing.T) {

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, []byte("12345"))
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h1 := h.Hash([]int64{1, 2, 3}).H
	h2 := h.Hash([]int64{2, 3, 4}).H

	if bytes.Equal(h1, h2) {
		t.Fatal("Slices the same, when should be different")
	}
}

func TestSlicesOfInterfaces(t *testing.T) {

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, []byte("12345"))
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h1 := h.Hash([]interface{}{1, 2, 3}).H
	h2 := h.Hash([]interface{}{1, 2, 3}).H

	if !bytes.Equal(h1, h2) {
		t.Fatal("Slices the same, hashes should not be different")
	}
}

func TestSlicesOfDifferentTypes(t *testing.T) {

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, []byte("12345"))
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h1 := h.Hash([]interface{}{1, "Hello", float64(3.22)}).H
	h2 := h.Hash([]interface{}{1, "Hello", float64(3.22)}).H

	if !bytes.Equal(h1, h2) {
		t.Fatal("Slices the same, hashes should not be different")
	}
}

func TestSlicesOfDifferentTypesAndOrders(t *testing.T) {

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, []byte("12345"))
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h1 := h.Hash([]interface{}{1, "Hello", float64(3.22)}).H
	h2 := h.Hash([]interface{}{"Hello", 1, float64(3.22)}).H

	if !bytes.Equal(h1, h2) {
		t.Fatal("Slices equivalent, hashes should not be different")
	}
}

func TestNestedSlicesOfDifferentTypesAndOrders(t *testing.T) {

	h, err := DefaultFactory.GetHasherWithSalt(SHA256, []byte("12345"))
	if err != nil {
		t.Fatal("Unable to retrieve SHA256 Hasher")
	}

	h1 := h.Hash([]interface{}{1, []interface{}{"Hello", "World", 2}, float64(3.22)}).H
	h2 := h.Hash([]interface{}{[]interface{}{"Hello", "World", 2}, 1, float64(3.22)}).H

	if !bytes.Equal(h1, h2) {
		t.Fatal("Slices equivalent, hashes should not be different")
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
