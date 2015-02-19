// Package finder enables searching slices of any struct type.
// Currently, it works for structs that implement a GetKey() string method
// thereby satisfying the internal hasKey interface.
// TODO: allow more generic searching by any field passed in

package finder

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// hasKey interface provKeyes an object to convert generic slices into for sorting and searching
type hasKey interface {
	GetKey() string
}

// type ByKey provKeyes a named type so []hasKey can be used as a receiver
type ByKey []hasKey

// hasKey implements sort
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].GetKey() < a[j].GetKey() }

// FindByKey takes a []hasKey (or &[]hasKey)to search and a string containing the Key.
// It will use sort.Search (binary search) and return the matching hasKey object
func FindByKey(items interface{}, Key string) (found hasKey, err error) {

	s := reflect.ValueOf(items).Elem()
	a := make(ByKey, s.Len())

	for i := 0; i < s.Len(); i++ {
		hasKey, ok := s.Index(i).Interface().(hasKey)
		if !ok {
			return nil, fmt.Errorf("Slice of type %T doesn't implement HasKey", s.Index(i).Interface())
		}
		a[i] = hasKey
	}

	sort.Sort(a)

	foundIndex := sort.Search(len(a), func(i int) bool {
		return a[i].GetKey() >= Key
	})
	if foundIndex == a.Len() || a[foundIndex].GetKey() != Key {
		return nil, errors.New("Key not found")
	}
	found = a[foundIndex]

	return found, nil
}
