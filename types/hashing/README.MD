## Hashing

The Hashing module provides a generalised mechanism to provide hashing 
across an arbitrary set of algorithms, with `SHA256` implemented directly.

Implementations of the `Hasher` interface generate hashes accordingly to
their specific algorithm, the type of which is retrievable via the `GetType`
function.

Instances of `Hasher` are created by their corresponding implementations of
the `HasherCreator` interface.

`HasherCreator` instances are not used directly, but instead are added to
an instance of `Factory` which is created by calling `NewFactory`.  The `Factory` type implements `HasherFactory`, which then provides direct access to `Hasher` instances based on the algorithm desired.

The `DefaultFactory` is available and provides access to the module's SHA256 implementation.

### SHA256 implementation

The module's internal SHA256 will provide a consistent hash for all
types.

This includes:
* across slices with the same content but differring element ordering
* looking through pointer values
* where slices or maps have instances of different types (e.g. `[]interface{}` or `map[string]interface{}`)
* arbitary depth of objects within slice or map

### Example

```
func main() {
    getHash := func(i interface{}, salt []byte) []byte {
        h, _ := DefaultFactory.GetHasherWithSalt(SHA256, salt)
        return h.Hash(i).H
    }


    fmt.Printf("%x\n", getHash([]int64{1, 2, 3, 4, 5}, []byte("12345")))
    // Output: cd5b761f3877aacab78715384f023546cbf462b1ece14de26f7e1ebad1ac591e
}
```