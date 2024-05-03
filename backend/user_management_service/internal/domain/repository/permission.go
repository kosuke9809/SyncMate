package repository

import (
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IPermissionRepository interface {
	Create(permission *model.Permission) (*model.Permission, error)
	Update(permission *model.Permission) (*model.Permission, error)
	Delete(id uint) error
	FindByID(id uint) (*model.Permission, error)
	FindByName(name string) (*model.Permission, error)
	GetAll() ([]*model.Permission, error)
}
