package models

import "time"

type User struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	DeletedOn time.Time `json:"deleted_on"`
}

// CreateUser create a User
func CreateUser(id uint, email, password string) *User {
	return &User{
		Id:        id,
		Email:     email,
		Password:  password,
		CreatedOn: time.Now(),
	}
}
