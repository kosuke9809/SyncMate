package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"github.com/kosuke9809/SyncMate/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(username, email, password string) (*model.User, error)
	Authenticate(email, password string) (*model.User, error)
	GeneratePasswordResetToken(email string) (string, error)
	ResetPassword(token, newPassword string) error
	FindByEmail(email string) (*model.User, error)
}

type userService struct {
	ur repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &userService{ur}
}

func (us *userService) CreateUser(username, email, password string) (*model.User, error) {
	existingUser, _ := us.ur.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:       uuid.New(),
		Username: username,
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

func (us *userService) GeneratePasswordResetToken(email string) (string, error) {
	user, err := us.ur.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	token, err := utils.GenerateToken(user.ID.String(), 1440)
	if err != nil {
		return "", err
	}
	user.PasswordResetToken = token
	user.PasswordResetExpires = time.Now().Add(24 * time.Hour)

	if err := us.ur.Update(user); err != nil {
		return "", err
	}
	return token, nil
}

func (us *userService) ResetPassword(token, newPassword string) error {
	user, err := us.ur.FindByPasswordResetToken(token)
	if err != nil {
		return err
	}
	if user == nil || user.PasswordResetExpires.Before(time.Now()) {
		return errors.New("invalid or expired token")
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hp)
	user.PasswordResetToken = ""
	user.PasswordResetExpires = time.Time{}

	if err := us.ur.Update(user); err != nil {
		return err
	}
	return nil
}

func (us *userService) FindByEmail(email string) (*model.User, error) {
	return us.ur.FindByEmail(email)
}
