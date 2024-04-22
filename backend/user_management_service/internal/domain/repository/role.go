package repository

import "github.com/kosuke9809/SyncMate/internal/domain/model"

type IRoleRepository interface {
	Create(role *model.Role) (*model.Role, error)
	Update(role *model.Role) (*model.Role, error)
	Delete(id int) error
	FindById(id int) (*model.Role, error)
	FindByName(name string) (*model.Role, error)
	GetAll() ([]*model.Role, error)
}
