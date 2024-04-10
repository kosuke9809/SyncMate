package repository

import "github.com/kosuke9809/SyncMate/internal/domain/model"

type RoleRepository interface {
	Create(role *model.Role) (*model.Role, error)
	Update(role *model.Role) (*model.Role, error)
	Delete(id int) error
	GetById(id int) (*model.Role, error)
	GetByName(name string) (*model.Role, error)
	GetAll() ([]*model.Role, error)
}

type RolePermissionRepository interface {
	AddPermissionToRole(roleID int, permissionID int) error
	RemovePermissionFromRole(roleID int, permissionID int) error
	GetRolePermissions(roleID int) ([]*model.Permission, error)
}
