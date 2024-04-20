package router

import (
	"github.com/kosuke9809/SyncMate/internal/presentation/http/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, uh handler.IUserHandler) {
	e.POST("/api/signup", uh.SignUp)
	e.POST("/api/signin", uh.SignIn)
	e.POST("/api/refresh", uh.RefreshAccessToken)
	e.POST("/api/request-password-reset", uh.RequestPasswordReset)
	e.POST("/api/reset-password", uh.ResetPassword)
}
