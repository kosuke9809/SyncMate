package repository

import "github.com/kosuke9809/SyncMate/internal/domain/model"

type IRolePermissionRepository interface {
	HasPermission(roleID, permissionID uint) (bool, error)
	AddPermissionToRole(roleID uint, permissionID uint) error
	RemovePermissionFromRole(roleID uint, permissionID uint) error
	GetRolePermissions(roleID uint) ([]*model.Permission, error)
}
