## Value

This module uses Protobufs to serialise values for multi-language transfers.

`NewValue` will create an instance `Value`  from an instance of `interface{}`, and `ParseValue` performs the reverse transformation to recover the `interface{}`.

Types supported:

* Basic types: bool, int32/64, uint32/64, float32/64, string
* `time.Time` and `time.Duration`
* Pointers of the basic types and the time types
* `map[string]interface{}`, where `interface{}` can be any of the basic types, pointers, slices or another map
* slices of the above types and `interface{}`, to an arbitrary depth (i.e `[][][][]int64 etc. are supported)

An error is returned if the type is not supported.

### Proto

The `Value` message uses `oneof`, ensuring that only a single value is included in each instance.

`ValueList` provides an list of `Value`, of length zero of more.

`ValueMap` provides a mapping `map<string, Value>`.

`Value` can contain a `ValueList` or a `ValueMap`, allowing complex structures to be modelled (and hence serialised) within a single `Value`. 

### Parsing slices

When parsing a `ValueList`, the `ValueListType` enumeration stored within it allows the parser to attempt to return a strongly typed slice, but will default to `[]interface{}` if this is not possible.


