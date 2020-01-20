package models

type Resource struct {
	ID          int
	Name        string
	Description string
	CreatedBy   *User
	Object      interface{}
}
