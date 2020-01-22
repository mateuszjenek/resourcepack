package models

// Resource represents reservable thing in system.
type Resource struct {
	ID          int
	Name        string
	Description string
	CreatedBy   *User
	Object      interface{}
}
