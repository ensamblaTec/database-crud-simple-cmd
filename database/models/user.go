package database

import (
	"fmt"
	"programa3/database"
	"programa3/internal/models"
	"time"
)

func CreateUser(user models.User) error {
	query := "insert into users (email,password,created_on) values ($1, $2, $3)"
	ok, err := database.GetConn().Exec(query, user.Email, user.Password, time.Now())
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func UpdateUser(user models.User) error {
	query := "update users set email=$1, password=$2, updated_on=$3, where id=$4"
	ok, err := database.GetConn().Exec(query, user.Email, user.Password, time.Now(), user.Id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func DeleteUser(id uint) error {
	query := "delete from users where id=$1"
	ok, err := database.GetConn().Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func SelectUser(id uint) error {
	query := "select * from users where id=$1"
	ok, err := database.GetConn().Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}
