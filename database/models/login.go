package database

import (
	"database/sql"
	"errors"
	"programa3/database"
)

func GetUserLogin(email, password string) error {
	conn := "select email, password from users where email=$1"
	ok := database.GetConn().QueryRow(conn, email)
	var resEmail, resPassword string
	err := ok.Scan(&resEmail, &resPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("cannot find results")
		} else {
			return err
		}
	}
	return nil
}
