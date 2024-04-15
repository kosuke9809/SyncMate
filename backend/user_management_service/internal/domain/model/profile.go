package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	IconImage  string `gorm:"type:varchar(255);"`
	CoverImage string `gorm:"type:varchar(255);"`
	Bio        string `gorm:"text"`
	Location   string `gorm:"type:varchar(100);"`
	Allergies  string `gorm:"type:text"`
	Preference string `gorm:"type:text"`
}