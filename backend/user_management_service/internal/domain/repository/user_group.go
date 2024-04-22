package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IUserGroupRepository interface {
	AddUserToGroup(userId uuid.UUID, groupId uuid.UUID, roleId uint) error
	RemoveUserFromGroup(userId uuid.UUID, groupId uuid.UUID) error
	GetGroupsByUserId(userId uuid.UUID) ([]*model.Group, error)
	FindUserGroupRole(userId uuid.UUID, groupId uuid.UUID) (*model.Role, error)
	GetMembersByGroupId(groupId uuid.UUID) ([]*model.User, error)
	FindGroupOwner(groupId uuid.UUID) (*model.User, error)
}
