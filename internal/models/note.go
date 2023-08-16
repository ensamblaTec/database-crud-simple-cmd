package models

import "time"

type Note struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []Tag     `json:"tags"`
	Owner       *User     `json:"owner"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	DeletedOn   time.Time `json:"deleted_on"`
}

// CreateNote create a Note
func CreateNote(id uint, title, description string) *Note {
	return &Note{
		Id:          id,
		Title:       title,
		Description: description,
	}
}
