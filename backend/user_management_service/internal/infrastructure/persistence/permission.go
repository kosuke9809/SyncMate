package persistence

import (
	"time"

	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type permissionPersistence struct {
	db *gorm.DB
}

func NewPermissionPersistence(db *gorm.DB) repository.IPermissionRepository {
	return &permissionPersistence{db}
}

func (pp *permissionPersistence) Create(permission *model.Permission) (*model.Permission, error) {
	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()
	if err := pp.db.Create(permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (pp *permissionPersistence) Update(permission *model.Permission) (*model.Permission, error) {
	permission.UpdatedAt = time.Now()
	if err := pp.db.Save(permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (pp *permissionPersistence) Delete(id uint) error {
	if err := pp.db.Delete(&model.Permission{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (pp *permissionPersistence) FindByID(id uint) (*model.Permission, error) {
	permission := &model.Permission{}
	err := pp.db.First(permission, id).Error
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (pp *permissionPersistence) FindByName(name string) (*model.Permission, error) {
	permission := &model.Permission{}
	err := pp.db.Where("permission_name = ?", name).First(permission).Error
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (pp *permissionPersistence) GetAll() ([]*model.Permission, error) {
	permissions := []*model.Permission{}
	if err := pp.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
