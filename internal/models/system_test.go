package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemIDFilter(t *testing.T) {
	testedSystem := &System{ID: 1}

	t.Run("Should return true", func(t *testing.T) {
		t.Run("When system has given ID", func(t *testing.T) {
			isMatched := SystemIDFilter(1)(testedSystem)
			assert.True(t, isMatched)
		})
	})

	t.Run("Should return false", func(t *testing.T) {
		t.Run("When system has not given ID", func(t *testing.T) {
			isMatched := SystemIDFilter(2)(testedSystem)
			assert.False(t, isMatched)
		})
	})
}

func TestSystemTypeFilter(t *testing.T) {
	testedSystem := &System{Type: SystemTypeSimulator}

	t.Run("Should return true", func(t *testing.T) {
		t.Run("When system is in given Type", func(t *testing.T) {
			isMatched := SystemTypeFilter(SystemTypeSimulator)(testedSystem)
			assert.True(t, isMatched)
		})
	})

	t.Run("Should return false", func(t *testing.T) {
		t.Run("When system is not in given Type", func(t *testing.T) {
			isMatched := SystemTypeFilter(SystemTypeHardware)(testedSystem)
			assert.False(t, isMatched)
		})
	})
}

func TestSystemAllFilter(t *testing.T) {
	testedSystem := &System{}

	t.Run("Should return always true", func(t *testing.T) {
		isMatched := SystemAllFilter()(testedSystem)
		assert.True(t, isMatched)
	})
}

func TestSystemCollectionMapToID(t *testing.T) {
	testedCollection := SystemCollection{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	}
	t.Run("Should return slice of systems' IDs from SystemCollection", func(t *testing.T) {
		expectedSystemIDs := []int{1, 2, 3, 4}
		systemIDs := testedCollection.MapToID()
		assert.EqualValues(t, expectedSystemIDs, systemIDs)
	})
}
