package models

type User struct {
	Email        string
	Username     string
	Passwordhash string
	Fullname     string
	CreateDate   string
	Role         int
}

// validate password
func (u *User) ValidatePasswordHash(pswdhash string) bool {
	return u.Passwordhash == pswdhash
}

// equality checking function
func (u *User) Equals(other User) bool {
	return (u.Email == other.Email ||
		u.Username == other.Username)
}
