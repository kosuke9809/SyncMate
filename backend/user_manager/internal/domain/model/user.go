package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username  string    `gorm:"type:varchar(30);not null"`
	Email     string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	ProfileID uint
	Profile   Profile `gorm:"foreignKey:ProfileID;"`
	Groups    []Group `gorm:"many2many:user_groups;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserGroup struct {
	UserID  uuid.UUID `gorm:"type:uuid;primary_key;"`
	GroupID uint      `gorm:"type:int;primary_key;"`
	User    User      `gorm:"foreignKey:UserID;"`
	Group   Group     `gorm:"foreignKey:GroupID;"`
	RoleID  uint
	Role    Role `gorm:"foreignKey:RoleID;"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

type UserDetailResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Profile  Profile   `json:"profile"`
}
