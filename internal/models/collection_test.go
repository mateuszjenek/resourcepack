//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "Generic=System,User,Reservation"

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericCollectionFilter(t *testing.T) {
	t.Run("Should return Generic that fulfills given condition", func(t *testing.T) {
		testedCollection := GenericCollection{
			new(Generic),
			new(Generic),
		}

		t.Run("When given condition match everything", func(t *testing.T) {
			fileredCollection := testedCollection.Filter(func(user *Generic) bool { return true })
			assert.EqualValues(t, testedCollection, fileredCollection)
		})

		t.Run("When given condition match nothing", func(t *testing.T) {
			fileredCollection := testedCollection.Filter(func(user *Generic) bool { return false })
			emptyCollection := GenericCollection{}
			assert.EqualValues(t, emptyCollection, fileredCollection)
		})
	})
}
