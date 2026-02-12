package models

type Post struct {
	BaseModel
	AuthorID uint   `gorm:"not null; index"`
	Title    string `gorm:"type:varchar(500); not null"`
	Text     string `gorm:"type:text; not null"`

	Author   User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}
