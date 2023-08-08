package models

type User struct {
	Id              uint
	Email, Password string
}

// CreateUser create a User
func CreateUser(id uint, email, password string) *User {
	return &User{
		id,
		email,
		password,
	}
}
