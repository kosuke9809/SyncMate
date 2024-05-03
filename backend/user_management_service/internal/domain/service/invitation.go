package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IInvitationService interface {
	SendInvitation(groupID, inviteeID, inviterID uuid.UUID) (*model.Invitation, error)
	AcceptInvitation(invitationID, inviteeID uuid.UUID) error
	RejectInvitation(invitationID, inviteeID uuid.UUID) error
	CancelInvitation(invitationID, inviterID uuid.UUID) error
	GetInvitationsByGroup(groupID uuid.UUID) ([]*model.Invitation, error)
	GetInvitationsByInvitee(inviteeID uuid.UUID) ([]*model.Invitation, error)
	GetInvitationsByInviter(inviterID uuid.UUID) ([]*model.Invitation, error)
	FindByID(invitationID uuid.UUID) (*model.Invitation, error)
}

type invitationService struct {
	ir repository.IInvitationRepository
	ur repository.IUserRepository
	gr repository.IGroupRepository
}

func NewInvitationService(ir repository.IInvitationRepository, ur repository.IUserRepository, gr repository.IGroupRepository) IInvitationService {
	return &invitationService{
		ir: ir,
		ur: ur,
		gr: gr,
	}
}

func (is *invitationService) SendInvitation(groupID, inviteeID, inviterID uuid.UUID) (*model.Invitation, error) {
	if _, err := is.gr.FindByID(groupID); err != nil {
		return nil, err
	}
	if _, err := is.ur.FindByID(inviteeID); err != nil {
		return nil, err
	}
	if _, err := is.ur.FindByID(inviterID); err != nil {
		return nil, err
	}
	invitation := &model.Invitation{
		GroupID:   groupID,
		InviteeID: inviteeID,
		InviterID: inviterID,
	}
	err := is.ir.Create(invitation)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (is *invitationService) AcceptInvitation(invitationID, inviteeID uuid.UUID) error {
	invitation, err := is.ir.FindByID(invitationID)
	if err != nil {
		return err
	}
	if invitation.InviteeID != inviteeID {
		return errors.New("invitation is not for this user")
	}
	if invitation.Status != "pending" {
		return errors.New("invitation is not pending")
	}
	invitation.Status = "accepted"
	if err := is.ir.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (is *invitationService) RejectInvitation(invitationID, inviteeID uuid.UUID) error {
	invitation, err := is.ir.FindByID(invitationID)
	if err != nil {
		return err
	}
	if invitation.InviteeID != inviteeID {
		return errors.New("invitation is not for this user")
	}
	if invitation.Status != "pending" {
		return errors.New("invitation is not pending")
	}
	invitation.Status = "rejected"
	if err := is.ir.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (is *invitationService) CancelInvitation(invitationID, inviterID uuid.UUID) error {
	invitation, err := is.ir.FindByID(invitationID)
	if err != nil {
		return err
	}
	if invitation.InviterID != inviterID {
		return errors.New("invitation is not for this user")
	}
	if invitation.Status != "pending" {
		return errors.New("invitation is not pending")
	}
	invitation.Status = "canceled"
	if err := is.ir.Update(invitation); err != nil {
		return err
	}
	return nil
}

func (is *invitationService) GetInvitationsByGroup(groupID uuid.UUID) ([]*model.Invitation, error) {
	if _, err := is.gr.FindByID(groupID); err != nil {
		return nil, err
	}
	invitations, err := is.ir.FindByGroupID(groupID)
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (is *invitationService) GetInvitationsByInvitee(inviteeID uuid.UUID) ([]*model.Invitation, error) {
	if _, err := is.ur.FindByID(inviteeID); err != nil {
		return nil, err
	}
	invitations, err := is.ir.FindByInviteeID(inviteeID)
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (is *invitationService) GetInvitationsByInviter(inviterID uuid.UUID) ([]*model.Invitation, error) {
	if _, err := is.ur.FindByID(inviterID); err != nil {
		return nil, err
	}
	invitations, err := is.ir.FindByInviterID(inviterID)
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (is *invitationService) FindByID(invitationID uuid.UUID) (*model.Invitation, error) {
	invitation, err := is.ir.FindByID(invitationID)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}
