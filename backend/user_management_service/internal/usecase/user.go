package usecase

import (
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/service"
	"github.com/kosuke9809/SyncMate/internal/utils"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	SignIn(user model.User) (string, string, error)
	RefreshAccessToken(refreshToken string) (string, error)
	RequestPasswordReset(email string) (string, error)
	ResetPassword(token, newPassword string) error
}

type userUsecase struct {
	us service.IUserService
}

func NewUserUsecase(us service.IUserService) *userUsecase {
	return &userUsecase{us}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	createdUser, err := uu.us.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		return model.UserResponse{}, err
	}
	res := model.UserResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}
	return res, nil
}

func (uu *userUsecase) SignIn(user model.User) (string, string, error) {
	authUser, err := uu.us.Authenticate(user.Email, user.Password)
	if err != nil {
		return "", "", err
	}
	accessToken, err := utils.GenerateToken(authUser.ID.String(), 30) // 30 minutes
	if err != nil {
		return "", "", err
	}
	refreshToken, err := utils.GenerateToken(authUser.ID.String(), 1440) // 24 hours
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (uu *userUsecase) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}
	newAccessToken, err := utils.GenerateToken(claims.UserID, 30)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}

func (uu *userUsecase) RequestPasswordReset(email string) (string, error) {
	token, err := uu.us.GeneratePasswordResetToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uu *userUsecase) ResetPassword(token, newPassword string) error {
	if err := uu.us.ResetPassword(token, newPassword); err != nil {
		return err
	}
	return nil
}
