package database

import (
	"fmt"
	"programa3/database"
	"programa3/internal/models"
	"time"
)

func CreateNote(note models.Note) error {
	conn := "insert into notes(title, description, owner, created_on) values($1,$2,$3,$4)"
	ok, err := database.GetConn().Exec(conn, note.Title, note.Description, note.Owner, time.Now())
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func DeleteNote(id uint) error {
	conn := "delete from notes where id=$1"
	ok, err := database.GetConn().Exec(conn, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func UpdateNote(note models.Note) error {
	conn := "update from notes set title=$1, description=$2, owner=$3, updated_on=$4"
	ok, err := database.GetConn().Exec(conn, note.Title, note.Description, note.Owner, time.Now())
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func SelectNote(id uint) error {
	conn := "select * from notes where id=$1"
	ok, err := database.GetConn().Exec(conn, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}
