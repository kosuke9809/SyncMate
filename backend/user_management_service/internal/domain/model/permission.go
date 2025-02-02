package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	PermissionName string  `gorm:"type:varchar(255);unique;not null"`
	Roles          []*Role `gorm:"many2many:role_permissions"`
}
