package database

import (
	"database/sql"
	"errors"
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

func GetNote(id uint) error {
	conn := "select * from notes where id=$1"
	ok, err := database.GetConn().Exec(conn, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func GetNotes(email string) ([]string, error) {
	conn := "select title, created_on from notes where owner=(select id_user from users where email=$1)"
	ok, err := database.GetConn().Query(conn, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("cannot find results")
		} else {
			return nil, err
		}
	}
	var records []string
	for ok.Next() {
		var title, created_on string
		if err := ok.Scan(&title, &created_on); err != nil {
			return nil, err
		}
		record := fmt.Sprintf("Titulo:\t%s, Field2:\t%s", title, created_on)
		records = append(records, record)
	}
	return records, nil
}
