package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IUserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
	FindByEmail(email string) (*model.User, error)
	FindById(id uuid.UUID) (*model.User, error)
	FindProfileById(userId uuid.UUID) (*model.Profile, error)
	GetAll() ([]*model.User, error)
}

type IUserGroupRepository interface {
	AddUserToGroup(userId uuid.UUID, groupId uuid.UUID, roleId uint) error
	RemoveUserFromGroup(userId uuid.UUID, groupId uuid.UUID) error
	GetGroupsByUserId(userId uuid.UUID) ([]*model.Group, error)
	GetGroupMembers(groupId uuid.UUID) ([]*model.User, error)
}
