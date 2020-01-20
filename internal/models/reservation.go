package models

import "time"

// Reservation represents single reservation entity
type Reservation struct {
	ID        int
	Author    string
	EndTime   time.Time
	StartTime time.Time
	SystemIDs []int
}

// ReservationActiveFilter returns filter to get only active reservations
func ReservationActiveFilter() ReservationCollectionFilter {
	return func(reservation *Reservation) bool {
		return time.Now().Before(reservation.EndTime)
	}
}

// ReservationIDFilter returns filter to get only reservations with given ID
func ReservationIDFilter(id int) ReservationCollectionFilter {
	return func(reservation *Reservation) bool {
		return reservation.ID == id
	}
}

// Builder for making reservations
type Builder struct {
	systems  SystemCollection
	User     User
	Duration time.Duration
}

// Reserve final step in building process, reserve given systems
func (b *Builder) Reserve() (*Reservation, error) {
	reservation := &Reservation{
		Author:    b.User.Username,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(b.Duration),
		SystemIDs: b.systems.MapToID(),
	}
	return reservation, nil
}

// FromRepository add systems to builder to prepare them for reserving
func (b *Builder) FromRepository(systems SystemCollection) error {
	panic("Not implemented")
}

// FromSimulatorPool add number of systems from public simulator pool to builder to prepare them for reserving
func (b *Builder) FromSimulatorPool(numberOfSimulators int) error {
	panic("Not implemented")
}

// FromHardware add number of hardware systems from repository to builder to prepare them for reserving
func (b *Builder) FromHardware(numberOfHardware int) error {
	panic("Not implemented")
}

// FromOpenStack add number of systems from user openstack to builder to prepare them for reserving
func (b *Builder) FromOpenStack(numberOfPrivateSimulators int) error {
	panic("Not implemented")
}
