package hashing

import "fmt"

func ExampleDefaultFactory() {
	getHash := func(i interface{}, salt []byte) []byte {
		h, _ := DefaultFactory.GetHasherWithSalt(SHA256, salt)
		return h.Hash(i).H
	}

	fmt.Printf("%x\n", getHash([]int64{1, 2, 3, 4, 5}, []byte("12345")))
	// Output: cd5b761f3877aacab78715384f023546cbf462b1ece14de26f7e1ebad1ac591e
}
