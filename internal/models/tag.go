package models

import "time"

type Tag struct {
	Id        uint      `json:"id"`
	Tag       string    `json:"tag"`
	Color     string    `json:"color"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	DeletedOn time.Time `json:"deleted_on"`
}

// CreateTag Create a Tag
func CreateTag(id uint, nameTag, color string) *Tag {
	return &Tag{
		Id:    id,
		Tag:   nameTag,
		Color: color,
	}
}
