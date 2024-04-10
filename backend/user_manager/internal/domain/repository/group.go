package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type GroupRepository interface {
	Create(group *model.Group) (*model.Group, error)
	Update(group *model.Group) (*model.Group, error)
	Delete(id uuid.UUID) error
	GetAll() ([]*model.Group, error)
	GetById(id uuid.UUID) (*model.Group, error)
	GetByName(name string) (*model.Group, error)
}
