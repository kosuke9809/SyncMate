package repository

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
)

type IUserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
	FindByEmail(email string) (*model.User, error)
	FindById(id uuid.UUID) (*model.User, error)
	FindProfileById(id uuid.UUID) (*model.Profile, error)
	FindByPasswordResetToken(token string) (*model.User, error)
	GetAll() ([]*model.User, error)
}
