package repository

import "github.com/kosuke9809/SyncMate/internal/domain/model"

type IProfileRepository interface {
	Create(profile *model.Profile) error
	Update(profile *model.Profile) error
	Delte(id uint) error
	FindByID(id uint) (*model.Profile, error)
	FindByUserID(userID uint) (*model.Profile, error)
}
