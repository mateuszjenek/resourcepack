package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReservationActiveFilter(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		t.Run("When reservation is active", func(t *testing.T) {
			testedReservation := &Reservation{EndTime: time.Now().Add(time.Hour)}
			isMatched := ReservationActiveFilter()(testedReservation)
			assert.True(t, isMatched)
		})
	})

	t.Run("Should return false", func(t *testing.T) {
		t.Run("When reservation is inactive", func(t *testing.T) {
			testedReservation := &Reservation{EndTime: time.Now().Add(-1 * time.Hour)}
			isMatched := ReservationActiveFilter()(testedReservation)
			assert.False(t, isMatched)
		})
	})
}

func TestReservationIDFilter(t *testing.T) {
	testedReservation := &Reservation{ID: 1}

	t.Run("Should return true", func(t *testing.T) {
		t.Run("When reservation has given ID", func(t *testing.T) {
			isMatched := ReservationIDFilter(1)(testedReservation)
			assert.True(t, isMatched)
		})
	})

	t.Run("Should return false", func(t *testing.T) {
		t.Run("When reservation has not given ID", func(t *testing.T) {
			isMatched := ReservationIDFilter(2)(testedReservation)
			assert.False(t, isMatched)
		})
	})
}
