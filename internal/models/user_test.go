package models

import "testing"

import "github.com/stretchr/testify/assert"

func TestUsernameFilter(t *testing.T) {
	testedUser := User{Username: "username"}
	t.Run("Should return true", func(t *testing.T) {
		t.Run("When user match condition", func(t *testing.T) {
			isMatched := UsernameFilter("username")(&testedUser)
			assert.True(t, isMatched)
		})
	})
	t.Run("Should return false", func(t *testing.T) {
		t.Run("When user don't match condition", func(t *testing.T) {
			isMatched := UsernameFilter("invalidUsername")(&testedUser)
			assert.False(t, isMatched)
		})
	})
}
