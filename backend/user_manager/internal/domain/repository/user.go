package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
	GetByEmail(email string) (*model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	GetAll() ([]*model.User, error)
	GetProfileByID(userID uuid.UUID) (*model.Profile, error)
}

type UserGroupRepository interface {
	GetGroupsByUserID(userID uuid.UUID) ([]*model.Group, error)
	AddUserToGroup(userID uuid.UUID, groupID uuid.UUID, roleID uint) error
	RemoveUserFromGroup(userID uuid.UUID, groupID uuid.UUID) error
	GetGroupMembers(groupID uuid.UUID) ([]*model.User, error)
}
