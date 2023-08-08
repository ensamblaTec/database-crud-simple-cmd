package models

type Tag struct {
	Id             uint
	NameTag, Color string
}

// CreateTag Create a Tag
func CreateTag(id uint, nameTag, color string) *Tag {
	return &Tag{
		Id:      id,
		NameTag: nameTag,
		Color:   color,
	}
}
