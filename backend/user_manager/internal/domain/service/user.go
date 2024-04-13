package service

import (
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(email, password string) (*model.User, error)
	Authenticate(email, password string) (*model.User, error)
}

type userService struct {
	ur repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) *userService {
	return &userService{ur}
}

func (us *userService) CreateUser(email, password string) (*model.User, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:    email,
		Password: string(hp),
	}
	if err := us.ur.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Authenticate(email, password string) (*model.User, error) {
	user, err := us.ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}
