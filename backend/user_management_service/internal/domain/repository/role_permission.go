package repository

import "github.com/kosuke9809/SyncMate/internal/domain/model"

type IRolePermissionRepository interface {
	AddPermissionToRole(roleId int, permissionId int) error
	RemovePermissionFromRole(roleId int, permissionId int) error
	GetRolePermissions(roleId int) ([]*model.Permission, error)
}
