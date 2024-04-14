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
}

type userUsecase struct {
	us service.IUserService
}

func NewUserCase(us service.IUserService) *userUsecase {
	return &userUsecase{us}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	createdUser, err := uu.us.CreateUser(user.Email, user.Password)
	if err != nil {
		return model.UserResponse{}, err
	}
	res := model.UserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
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
