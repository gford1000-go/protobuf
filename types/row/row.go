package row

import "errors"

type RowID int64
type AttributeName string

var errUnknownAttributeName = errors.New("unknown attribute name")
var errRowIDMismatch = errors.New("mismatch in RowIDs")
var errDuplicateAttribute = errors.New("duplicate attribute value")

// newRow returns a new instance with the specified rowID
func newRow(rowID RowID) *Row {
	return &Row{
		rowId: rowID,
		atts:  map[AttributeName]interface{}{},
	}
}

type Row struct {
	rowId RowID
	atts  map[AttributeName]interface{}
}

func (r *Row) GetID() RowID {
	return r.rowId
}

// GetAttributeNames returns the list of attribute names
// in this Row instance
func (r *Row) GetAttributeNames() []AttributeName {
	keys := make([]AttributeName, 0, len(r.atts))
	for k := range r.atts {
		keys = append(keys, k)
	}
	return keys
}

// Get returns the value of the specified attribute
func (r *Row) Get(attributeName AttributeName) (interface{}, error) {
	i, ok := r.atts[attributeName]
	if !ok {
		return nil, errUnknownAttributeName
	}
	return i, nil
}

// GetAll returns the entire map of name -> value
func (r *Row) GetAll() map[AttributeName]interface{} {
	return r.atts
}

// addAttribute populates the instances attribuute map,
// raising an error if an attempt is made to update an
// existing attribute
func (r *Row) addAttribute(attributeName AttributeName, v interface{}) error {
	_, ok := r.atts[attributeName]
	if ok {
		return errDuplicateAttribute
	}
	r.atts[attributeName] = v
	return nil
}

// Join returns a new Row instance that brings together all the
// attributes for input rows that match this instance's RowID.
// This instance and input rows are unchanged by this operation,
// however this is a shallow copy of the attributes from those rows.
//
// If ignoreOnRowIDMismatch is true, then mismatches on RowID are
// not reported as an error, otherwise processing stops immediately.
// If an attribute is duplicated, then processing stops immediately
func (r *Row) Join(ignoreOnRowIDMismatch bool, rows ...*Row) (*Row, error) {
	newRow := newRow(r.rowId)

	f := func(src *Row) error {
		if src.rowId != newRow.rowId {
			if ignoreOnRowIDMismatch {
				return nil
			}
			return errRowIDMismatch
		}
		for k, v := range src.atts {
			err := r.addAttribute(k, v)
			if err != nil {
				return err
			}
			newRow.atts[k] = v
		}

		return nil
	}

	err := f(r)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		err := f(row)
		if err != nil {
			return nil, err
		}
	}
	return newRow, nil
}
