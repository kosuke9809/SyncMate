package model

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	GroupName   string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Users       []*User   `gorm:"many2many:group_users;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}