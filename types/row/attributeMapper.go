package row

// AttributeIdentifier is the on the wire identifier of the
// attribute that a Cell represents
type AttributeIdentifier int32

// AttributeFromIdentifier provides a mapping from the
// on the wire identifier to the attribute name
type AttributeFromIdentifier interface {
	FindName(i AttributeIdentifier) (AttributeName, error)
}

// AttributeFromName provides a mapping from the attribute
// nae to the identifier value to be used for serialisation
type AttributeFromName interface {
	FindIdentifier(n AttributeName) (AttributeIdentifier, error)
}
