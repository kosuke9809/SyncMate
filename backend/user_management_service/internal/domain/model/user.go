package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Username             string    `gorm:"type:varchar(30);not null" json:"username"`
	Email                string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password             string    `gorm:"type:varchar(255);" json:"password"`
	Groups               []Group   `gorm:"many2many:user_groups;"`
	PasswordResetToken   string    `gorm:"type:varchar(255);" json:"password_reset_token"`
	PasswordResetExpires time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
