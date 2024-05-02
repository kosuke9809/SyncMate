package persistence

import (
	"time"

	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type rolePersistence struct {
	db *gorm.DB
}

func NewRolePersistence(db *gorm.DB) repository.IRoleRepository {
	return &rolePersistence{db}
}

func (rp rolePersistence) Create(role *model.Role) (*model.Role, error) {
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	if err := rp.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (rp rolePersistence) Update(role *model.Role) (*model.Role, error) {
	role.UpdatedAt = time.Now()
	if err := rp.db.Save(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (rp rolePersistence) Delete(roleID uint) error {
	if err := rp.db.Delete(&model.Role{}, roleID).Error; err != nil {
		return err
	}
	return nil
}

func (rp rolePersistence) FindByID(roleID uint) (*model.Role, error) {
	var role *model.Role
	if err := rp.db.First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (rp rolePersistence) FindByName(roleName string) (*model.Role, error) {
	var role *model.Role
	if err := rp.db.Where("role_name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (rp rolePersistence) GetAll() ([]*model.Role, error) {
	var roles []*model.Role
	if err := rp.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
