package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IGroupRepository interface {
	Create(group *model.Group) (*model.Group, error)
	Update(group *model.Group) (*model.Group, error)
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*model.Group, error)
	FindByName(name string) (*model.Group, error)
	GetAll() ([]*model.Group, error)
}

type IGroupUserRepository interface {
	GetMembersByGroupId(groupId uuid.UUID) ([]*model.User, error)
	FindGroupOwner(groupId uuid.UUID) (*model.User, error)
}
