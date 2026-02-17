package models

import "time"

type RefreshToken struct {
	BaseModel
	UserID    uint
	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
