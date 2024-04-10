package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string        `gorm:"type:varchar(255)"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
	gorm.Model
	PermissionName string  `gorm:"type:varchar(255)"`
	Roles          []*Role `gorm:"many2many:role_permissions"`
}
