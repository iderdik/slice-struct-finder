// Package finder enables searching slices of any struct type.
// Currently, it works for structs that implement a GetID() string method
// thereby satisfying the internal hasID interface.
// TODO: allow more generic searching by any field passed in

package finder

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// hasID interface provides an object to convert generic slices into for sorting and searching
type hasID interface {
	GetID() string
}

// type ByID provides a named type so []hasID can be used as a receiver
type ByID []hasID

// hasID implements sort
func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].GetID() < a[j].GetID() }

// FindByID takes a []hasID (or &[]hasID)to search and a string containing the ID.
// It will use sort.Search (binary search) and return the matching hasID object
func FindByID(items interface{}, id string) (found hasID, err error) {

	s := reflect.ValueOf(items).Elem()
	a := make(ByID, s.Len())

	for i := 0; i < s.Len(); i++ {
		hasid, ok := s.Index(i).Interface().(hasID)
		if !ok {
			return nil, fmt.Errorf("Slice of type %T doesn't implement HasID", s.Index(i).Interface())
		}
		a[i] = hasid
	}

	sort.Sort(a)

	foundIndex := sort.Search(len(a), func(i int) bool {
		return a[i].GetID() >= id
	})
	if foundIndex == a.Len() || a[foundIndex].GetID() != id {
		return nil, errors.New("ID not found")
	}
	found = a[foundIndex]

	return found, nil
}
