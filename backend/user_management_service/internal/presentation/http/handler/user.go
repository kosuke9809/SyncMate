package handler

import (
	"net/http"
	"time"

	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/usecase"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	SignUp(ctx echo.Context) error
	SignIn(ctx echo.Context) error
	RefreshAccessToken(ctx echo.Context) error
	RequestPasswordReset(ctx echo.Context) error
	ResetPassword(ctx echo.Context) error
}

type userHandler struct {
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) SignUp(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	res, err := uh.uu.SignUp(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to sign up"})
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (uh *userHandler) SignIn(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	accessToken, refreshToken, err := uh.uu.SignIn(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to sign in"})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "AccessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	ctx.SetCookie(&http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Signed in successfully"})
}

func (uh *userHandler) RefreshAccessToken(ctx echo.Context) error {
	refreshToken, err := ctx.Cookie("RefreshToken")
	if err != nil || refreshToken.Value == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Refresh token is required"})
	}
	newAccessToken, err := uh.uu.RefreshAccessToken(refreshToken.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to refresh access token"})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "AccessToken",
		Value:    newAccessToken,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Access token refreshed successfully"})
}

func (uh *userHandler) RequestPasswordReset(ctx echo.Context) error {
	var req struct {
		Email string `json:"email"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	token, err := uh.uu.RequestPasswordReset(req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to request password reset"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func (uh *userHandler) ResetPassword(ctx echo.Context) error {
	var req struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	if err := uh.uu.ResetPassword(req.Token, req.NewPassword); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to reset password"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Password reset successfully"})
}
