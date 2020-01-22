package models

import "time"

// Reservation represents single reservation entity
type Reservation struct {
	ID        int
	Author    string
	EndTime   time.Time
	StartTime time.Time
	Resources []*Resource
}
