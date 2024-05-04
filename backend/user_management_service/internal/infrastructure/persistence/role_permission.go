package persistence

import (
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type rolePermissionPersistence struct {
	db *gorm.DB
}

func NewRolePermissionPersistence(db *gorm.DB) repository.IRolePermissionRepository {
	return &rolePermissionPersistence{db}
}

func (rpp *rolePermissionPersistence) HasPermission(roleID, permissionID uint) (bool, error) {
	var c int64
	err := rpp.db.Table("role_permissions").Where("role_id = ? AND permission_id = ?", roleID, permissionID).Count(&c).Error
	if err != nil {
		return false, err
	}
	return c > 0, nil
}

func (rpp *rolePermissionPersistence) AddPermissionToRole(roleID uint, permissionID uint) error {
	role := &model.Role{}
	if err := rpp.db.First(role, roleID).Error; err != nil {
		return err
	}
	permission := &model.Permission{}
	if err := rpp.db.First(permission, permissionID).Error; err != nil {
		return err
	}

	if err := rpp.db.Model(role).Association("Permissions").Append(permission); err != nil {
		return err
	}
	return nil
}

func (rpp *rolePermissionPersistence) RemovePermissionFromRole(roleID uint, permissionID uint) error {
	role := &model.Role{}
	if err := rpp.db.First(role, roleID).Error; err != nil {
		return err
	}
	permission := &model.Permission{}
	if err := rpp.db.First(permission, permissionID).Error; err != nil {
		return err
	}

	if err := rpp.db.Model(role).Association("Permissions").Delete(permission); err != nil {
		return err
	}
	return nil
}

func (rpp *rolePermissionPersistence) GetRolePermissions(roleID uint) ([]*model.Permission, error) {
	role := &model.Role{}
	if err := rpp.db.First(role, roleID).Error; err != nil {
		return nil, err
	}
	permissions := []*model.Permission{}
	if err := rpp.db.Model(role).Association("Permissions").Find(&permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}
