package models

type Note struct {
	Id                 uint
	Title, Description string
	Tags               []Tag
	Owner              *User
}

// CreateNote create a Note
func CreateNote(id uint, title, description string) *Note {
	return &Note{
		Id:          id,
		Title:       title,
		Description: description,
	}
}
