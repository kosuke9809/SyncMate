package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string        `gorm:"type:varchar(255)"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
}
