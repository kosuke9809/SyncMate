package model

import (
	"time"

	"github.com/google/uuid"
)

type Invitation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;)"`
	GroupID   uuid.UUID `gorm:"type:uuid;not null"`
	InviterID uuid.UUID `gorm:"type:uuid;not null"`
	InviteeID uuid.UUID `gorm:"type:uuid;not null"`
	Status    string    `gorm:"type:varchar(255);not null;default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
