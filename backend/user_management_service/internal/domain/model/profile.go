package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID         uuid.UUID `gorm:"primaryKey;"`
	IconImage  string    `gorm:"type:varchar(255);"`
	CoverImage string    `gorm:"type:varchar(255);"`
	Bio        string    `gorm:"text"`
	Location   string    `gorm:"type:varchar(100);"`
	Allergies  string    `gorm:"type:text"`
	Preference string    `gorm:"type:text"`
	UserID     uuid.UUID
	User       User `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
