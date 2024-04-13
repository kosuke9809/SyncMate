package repository

import (
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type PermissionRepository interface {
	Create(permission *model.Permission) (*model.Permission, error)
	Update(permission *model.Permission) (*model.Permission, error)
	Delete(id uint) error
	FindById(id uint) (*model.Permission, error)
	GetAll() ([]*model.Permission, error)
}
