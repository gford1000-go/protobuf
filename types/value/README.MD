## Value

This module uses Protobufs to serialise values for multi-language transfers.

`NewValue` will create an instance `Value`  from an instance of `interface{}`, and `ParseValue` performs the reverse transformation to recover the `interface{}`.

Types supported:

* Basic types: bool, int32/64, uint32/64, float32/64, string
* `time.Time` and `time.Duration`
* Pointers of the basic types and the time types
* `map[string]interface{}`, where 
* slices of the above types and `interface{}`

An error is returned if the type is not supported.

