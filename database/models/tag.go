package database

import (
	"fmt"
	"programa3/database"
	"programa3/internal/models"
	"time"
)

func CreateTag(tag models.Tag) error {
	conn := "insert into tags(tag, color, created_on) values($1,$2,$3)"
	ok, err := database.GetConn().Exec(conn, tag.Tag, tag.Color, time.Now())
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func UpdateTag(tag models.Tag) error {
	conn := "update from tags set tag=$1, color=$2, updated_on=$3 where id=$4"
	ok, err := database.GetConn().Exec(conn, tag.Tag, tag.Color, time.Now(), tag.Id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func DeleteTag(id uint) error {
	conn := "delete from tags where id=$1"
	ok, err := database.GetConn().Exec(conn, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}

func SelectTag(id uint) error {
	conn := "select * from tags where id=$1"
	ok, err := database.GetConn().Exec(conn, id)
	if err != nil {
		return err
	}
	fmt.Println(ok.RowsAffected())
	return nil
}
