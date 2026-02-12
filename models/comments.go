package models

type Comment struct {
	BaseModel
	AuthorID uint   `gorm:"not null;index"`
	PostID   uint   `gorm:"not null;index"`
	Text     string `gorm:"type:text;not null"`
	Post     Post   `gorm:"foreignKey:PostID"`
	Author   User   `gorm:"foreignKey:AuthorID"`
}
