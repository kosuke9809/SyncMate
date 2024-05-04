package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/service"
)

type IGroupUsecase interface {
	CreateNewGroup(groupName, description string, creatorID uuid.UUID) (*model.Group, error)
	InviteUserToGroup(inviterID, groupID uuid.UUID, inviteeEmail string) (*model.Invitation, error)
	AcceptInvitation(invitationID, inviteeID uuid.UUID) error
	RejectInvitation(invitationID, inviteeID uuid.UUID) error
	CancelInvitation(invitationID, inviterID uuid.UUID) error
	RemoveUserFromGroup(userID, removeUserID, groupID uuid.UUID) error
	GetGroupDetails(groupID uuid.UUID) (*model.Group, error)
	GetGroupMembers(groupID uuid.UUID) ([]*model.User, error)
	DeleteGroup(groupID uuid.UUID) error
}

type groupUsecase struct {
	gs  service.IGroupService
	us  service.IUserService
	ugs service.IUserGroupService
	rs  service.IRoleService
	rps service.IRolePermissionService
	is  service.IInvitationService
}

func NewGroupUsecase(gs service.IGroupService, us service.IUserService, ugs service.IUserGroupService, rs service.IRoleService, rps service.IRolePermissionService, is service.IInvitationService) IGroupUsecase {
	return &groupUsecase{
		gs:  gs,
		us:  us,
		ugs: ugs,
		rs:  rs,
		rps: rps,
		is:  is,
	}
}

func (gu *groupUsecase) CreateNewGroup(groupName, description string, creatorID uuid.UUID) (*model.Group, error) {
	group, err := gu.gs.CreateGroup(groupName, description)
	if err != nil {
		return nil, err
	}

	role, err := gu.rs.FindByName("Owner")
	if err != nil {
		return nil, err
	}

	err = gu.ugs.AddUserToGroup(creatorID, group.ID, role.ID)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (gu *groupUsecase) InviteUserToGroup(inviterID, groupID uuid.UUID, inviteeEmail string) (*model.Invitation, error) {
	// email -> user
	invitee, err := gu.us.FindByEmail(inviteeEmail)
	if err != nil {
		return nil, errors.New("user not found")
	}
	// check inviter is a member of the group
	isMember, err := gu.ugs.IsUserInGroup(invitee.ID, groupID)
	if err != nil {
		return nil, errors.New("failed to check if inviter is a member of the group")
	}
	if isMember {
		return nil, errors.New("invitee is already a member of the group")
	}
	// check inviter role in the group
	inviterRole, err := gu.ugs.FindUserRoleInGroup(inviterID, groupID)
	if err != nil {
		return nil, errors.New("failed to check inviter role in the group")
	}
	// check inviter has permission to invite users to the group
	ok, err := gu.rps.HasPermission(inviterRole.ID, "ManageGroupMembers")
	if err != nil {
		return nil, errors.New("failed to check if inviter has permission to invite users to the group")
	}

	if !ok {
		return nil, errors.New("inviter does not have permission to invite users to the group")
	}
	// check invitee is not a member of the group
	isMember, err = gu.ugs.IsUserInGroup(invitee.ID, groupID)
	if err != nil {
		return nil, errors.New("failed to check if invitee is a member of the group")
	}
	if isMember {
		return nil, errors.New("invitee is already a member of the group")
	}
	// create invitation
	invitation, err := gu.is.SendInvitation(groupID, invitee.ID, inviterID)
	if err != nil {
		return nil, errors.New("failed to send invitation")
	}
	return invitation, nil
}

func (gu *groupUsecase) AcceptInvitation(invitationID, inviteeID uuid.UUID) error {
	// find invitation
	invitation, err := gu.is.FindByID(invitationID)
	if err != nil {
		return err
	}
	// check invitation status
	if invitation.Status != "pending" {
		return errors.New("invitation is not pending")
	}
	// accept invitation
	if err = gu.is.AcceptInvitation(invitationID, inviteeID); err != nil {
		return err
	}
	// member role
	role, err := gu.rs.FindByName("Member")
	if err != nil {
		return err
	}
	// add user to group
	if err = gu.ugs.AddUserToGroup(invitation.InviteeID, invitation.GroupID, role.ID); err != nil {
		return err
	}
	return nil
}

func (gu *groupUsecase) RejectInvitation(invitationID, inviteeID uuid.UUID) error {
	// find invitation
	invitation, err := gu.is.FindByID(invitationID)
	if err != nil {
		return err
	}
	// check invitation status
	if invitation.Status != "pending" {
		return errors.New("invitation is not pending")
	}
	// reject invitation
	if err = gu.is.RejectInvitation(invitationID, inviteeID); err != nil {
		return err
	}
	return nil
}

func (gu *groupUsecase) CancelInvitation(invitationID, inviterID uuid.UUID) error {
	// find invitation
	invitation, err := gu.is.FindByID(invitationID)
	if err != nil {
		return err
	}
	// check invitation status
	if invitation.Status != "pending" {
		return errors.New("invitation is not pending")
	}
	// cancel invitation
	if err = gu.is.CancelInvitation(invitationID, inviterID); err != nil {
		return err
	}
	return nil
}

func (gu *groupUsecase) RemoveUserFromGroup(userID, removeUserID, groupID uuid.UUID) error {
	// check user is a member of the group
	isMember, err := gu.ugs.IsUserInGroup(removeUserID, groupID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("user is not a member of the group")
	}
	// check user role in the group
	userRole, err := gu.ugs.FindUserRoleInGroup(userID, groupID)
	if err != nil {
		return err
	}
	// check user has permission to remove users from the group
	ok, err := gu.rps.HasPermission(userRole.ID, "ManageGroupMembers")
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user does not have permission to remove users from the group")
	}
	// remove user from group
	if err = gu.ugs.RemoveUserFromGroup(removeUserID, groupID); err != nil {
		return err
	}
	return nil
}

func (gu *groupUsecase) GetGroupDetails(groupID uuid.UUID) (*model.Group, error) {
	group, err := gu.gs.FindByID(groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (gu *groupUsecase) GetGroupMembers(groupID uuid.UUID) ([]*model.User, error) {
	members, err := gu.ugs.GetMembersByGroupID(groupID)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (gu *groupUsecase) DeleteGroup(groupID uuid.UUID) error {
	// check group exists
	_, err := gu.gs.FindByID(groupID)
	if err != nil {
		return err
	}
	// check group has no members
	members, err := gu.ugs.GetMembersByGroupID(groupID)
	if err != nil {
		return err
	}
	if len(members) > 0 {
		return errors.New("group has members")
	}
	// delete group
	if err = gu.gs.DeleteGroup(groupID); err != nil {
		return err
	}
	return nil
}
