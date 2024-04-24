package router

import (
	_ "github.com/kosuke9809/SyncMate/docs"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, uh handler.IUserHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/signup", uh.SignUp)
	e.POST("/api/signin", uh.SignIn)
	e.POST("/api/refresh", uh.RefreshAccessToken)
	e.POST("/api/request-password-reset", uh.RequestPasswordReset)
	e.POST("/api/reset-password", uh.ResetPassword)
}
