package models

type User struct {
	Email    string
	Username string
	Password string
}

// equality checking function
func (u *User) Equals(other User) bool {
	return (u.Email == other.Email ||
		u.Username == other.Username)
}
