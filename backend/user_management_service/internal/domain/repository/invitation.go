package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IInvitationRepository interface {
	Create(invitation *model.Invitation) error
	Update(invitation *model.Invitation) error
	FindByID(id uuid.UUID) (*model.Invitation, error)
	FindByGroupID(groupID uuid.UUID) ([]*model.Invitation, error)
	FindByInviteeID(inviteeID uuid.UUID) ([]*model.Invitation, error)
	FindByInviterID(inviterID uuid.UUID) ([]*model.Invitation, error)
	Delete(id uuid.UUID) error
}
