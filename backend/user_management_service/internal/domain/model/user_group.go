package model

import "github.com/google/uuid"

type UserGroup struct {
	UserID  uuid.UUID `gorm:"type:uuid;primary_key;"`
	GroupID uuid.UUID `gorm:"type:int;primary_key;"`
	RoleID  uint
	User    User  `gorm:"foreignKey:UserID;"`
	Group   Group `gorm:"foreignKey:GroupID;"`
	Role    Role  `gorm:"foreignKey:RoleID;"`
}
