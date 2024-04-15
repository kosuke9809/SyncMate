package persistence

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.IUserRepository {
	return &userPersistence{db}
}

func (up *userPersistence) Create(user *model.User) error {
	if err := up.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) Update(user *model.User) error {
	if err := up.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) Delete(id uuid.UUID) error {
	if err := up.db.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := up.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (up *userPersistence) FindById(id uuid.UUID) (*model.User, error) {
	user := model.User{}
	if err := up.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (up *userPersistence) FindProfileById(userId uuid.UUID) (*model.Profile, error) {
	profile := model.Profile{}
	if err := up.db.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (up *userPersistence) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := up.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
