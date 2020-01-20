package models

// User represent logged user
type User struct {
	Username string
	Email    string
	PassHash string
	Role     UserRole
}

// UserRole define user's privileges
type UserRole string

const (
	// UserRoleRegular is a default user role
	UserRoleRegular UserRole = "Regular"
	// UserRoleAdmin is a role with high privileges
	UserRoleAdmin UserRole = "Admin"
	// UserRoleRoot is an administrator of whole system
	UserRoleRoot UserRole = "Root"
)

// UsernameFilter returns filter to get only users with matching username
func UsernameFilter(username string) UserCollectionFilter {
	return func(user *User) bool {
		return user.Username == username
	}
}
