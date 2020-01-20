//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "Generic=System,User,Reservation"

package models

import "github.com/cheekybits/genny/generic"

// Generic is a type of collection
type Generic generic.Type

// GenericCollection represents generic items collection
type GenericCollection []*Generic

// Filter returns items which meet the given condition
func (collection GenericCollection) Filter(condition GenericCollectionFilter) GenericCollection {
	filteredCollection := GenericCollection{}
	for _, item := range collection {
		if condition(item) {
			filteredCollection = append(filteredCollection, item)
		}
	}
	return filteredCollection
}

// GenericCollectionFilter is a condition function for filtering items
type GenericCollectionFilter func(*Generic) bool
