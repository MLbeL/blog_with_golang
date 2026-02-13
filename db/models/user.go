package models

type User struct {
	BaseModel
	Email        string `gorm:"type:varchar(80); not null; uniqueIndex"`
	PasswordHash string `gorm:"type:varchar(255); not null"`
	Name         string `gorm:"type:varchar(255); not null"`

	Posts    []Post    `gorm:"foreignKey:AuthorID"`
	Comments []Comment `gorm:"foreignKey:AuthorID"`
}
