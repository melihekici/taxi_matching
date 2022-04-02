package models

// swagger:model User
type User struct {
	// User email
	// in: string
	Email string
	// Username
	// in: string
	Username string
	// Password
	// in: string
	Password string
}

// equality checking function
func (u *User) Equals(other User) bool {
	return (u.Email == other.Email ||
		u.Username == other.Username)
}
