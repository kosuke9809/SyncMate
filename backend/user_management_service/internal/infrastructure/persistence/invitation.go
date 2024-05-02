package persistence

import (
	"time"

	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type invitationPersistence struct {
	db *gorm.DB
}

func NewInvitationPersistence(db *gorm.DB) repository.IInvitationRepository {
	return &invitationPersistence{db}
}

func (ip *invitationPersistence) Create(invitation *model.Invitation) error {
	invitation.CreatedAt = time.Now()
	invitation.UpdatedAt = time.Now()
	err := ip.db.Create(invitation).Error
	if err != nil {
		return err
	}
	return nil
}

func (ip *invitationPersistence) Update(invitation *model.Invitation) error {
	invitation.UpdatedAt = time.Now()
	err := ip.db.Save(invitation).Error
	if err != nil {
		return err
	}
	return nil
}

func (ip *invitationPersistence) FindByID(id uuid.UUID) (*model.Invitation, error) {
	invitation := &model.Invitation{}
	err := ip.db.First(invitation, id).Error
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (ip *invitationPersistence) FindByGroupID(groupID uuid.UUID) ([]*model.Invitation, error) {
	invitations := []*model.Invitation{}
	err := ip.db.Where("group_id = ?", groupID).Find(&invitations).Error
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (ip *invitationPersistence) FindByInviteeID(inviteeID uuid.UUID) ([]*model.Invitation, error) {
	invitations := []*model.Invitation{}
	err := ip.db.Where("invitee_id = ?", inviteeID).Find(&invitations).Error
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (ip *invitationPersistence) FindByInviterID(inviterID uuid.UUID) ([]*model.Invitation, error) {
	invitations := []*model.Invitation{}
	err := ip.db.Where("inviter_id = ?", inviterID).Find(&invitations).Error
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (ip *invitationPersistence) Delete(id uuid.UUID) error {
	err := ip.db.Delete(&model.Invitation{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
