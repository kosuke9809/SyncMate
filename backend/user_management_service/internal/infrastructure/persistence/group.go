package persistence

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type groupPersistence struct {
	db *gorm.DB
}

func NewGroupPersistence(db *gorm.DB) repository.IGroupRepository {
	return &groupPersistence{db}
}

func (gp *groupPersistence) Create(group *model.Group) (*model.Group, error) {
	if err := gp.db.Create(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (gp *groupPersistence) Update(group *model.Group) (*model.Group, error) {
	if err := gp.db.Save(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (gp *groupPersistence) Delete(id uuid.UUID) error {
	if err := gp.db.Delete(&model.Group{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (gp *groupPersistence) FindByID(id uuid.UUID) (*model.Group, error) {
	group := &model.Group{}
	if err := gp.db.First(group, id).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (gp *groupPersistence) FindByName(name string) (*model.Group, error) {
	group := &model.Group{}
	if err := gp.db.Where("group_name = ?", name).First(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (gp *groupPersistence) GetAll() ([]*model.Group, error) {
	var groups []*model.Group
	if err := gp.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
