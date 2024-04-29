package service

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IUserGroupService interface {
	AddUserToGroup(userID uuid.UUID, groupID uuid.UUID, roleID uint) error
	RemoveUserFromGroup(userID uuid.UUID, groupID uuid.UUID) error
	GetGroupsByUserID(userID uuid.UUID) ([]*model.Group, error)
	FindUserGroupRole(userID uuid.UUID, groupID uuid.UUID) (*model.Role, error)
	GetMembersByGroupID(groupID uuid.UUID) ([]*model.User, error)
	FindGroupOwner(groupID uuid.UUID) (*model.User, error)
}

type userGroupService struct {
	ugr repository.IUserGroupRepository
}

func NewUserGroupService(ugr repository.IUserGroupRepository) IUserGroupService {
	return &userGroupService{ugr}
}

func (ugs *userGroupService) AddUserToGroup(userID uuid.UUID, groupID uuid.UUID, roleID uint) error {
	return ugs.ugr.AddUserToGroup(userID, groupID, roleID)
}

func (ugs *userGroupService) RemoveUserFromGroup(userID uuid.UUID, groupID uuid.UUID) error {
	return ugs.ugr.RemoveUserFromGroup(userID, groupID)
}

func (ugs *userGroupService) GetGroupsByUserID(userID uuid.UUID) ([]*model.Group, error) {
	return ugs.ugr.GetGroupsByUserID(userID)
}

func (ugs *userGroupService) FindUserGroupRole(userID uuid.UUID, groupID uuid.UUID) (*model.Role, error) {
	return ugs.ugr.FindUserGroupRole(userID, groupID)
}

func (ugs *userGroupService) GetMembersByGroupID(groupID uuid.UUID) ([]*model.User, error) {
	return ugs.ugr.GetMembersByGroupID(groupID)
}

func (ugs *userGroupService) FindGroupOwner(groupID uuid.UUID) (*model.User, error) {
	return ugs.ugr.FindGroupOwner(groupID)
}
