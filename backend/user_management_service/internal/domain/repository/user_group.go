package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IUserGroupRepository interface {
	AddUserToGroup(userID uuid.UUID, groupID uuid.UUID, roleID uint) error
	RemoveUserFromGroup(userID uuid.UUID, groupID uuid.UUID) error
	GetGroupsByUserID(userID uuid.UUID) ([]*model.Group, error)
	FindUserRoleInGroup(userID uuid.UUID, groupID uuid.UUID) (*model.Role, error)
	GetMembersByGroupID(groupID uuid.UUID) ([]*model.User, error)
	FindGroupOwner(groupID uuid.UUID) (*model.User, error)
	IsUserInGroup(userID uuid.UUID, groupID uuid.UUID) (bool, error)
}
