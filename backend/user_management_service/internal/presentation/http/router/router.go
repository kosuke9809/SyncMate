package router

import (
	_ "github.com/kosuke9809/SyncMate/docs"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/handler"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, uh handler.IUserHandler, gh handler.IGroupHandler) {
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// api group
	all := e.Group("/api")
	auth := e.Group("/api")
	middleware.RequireJWTAuth(auth)
	// user
	all.POST("/signup", uh.SignUp)
	all.POST("/signin", uh.SignIn)
	all.POST("/refresh", uh.RefreshAccessToken)
	auth.POST("/password/reset/request", uh.RequestPasswordReset)
	auth.POST("/password/reset", uh.ResetPassword)
	// group
	auth.POST("/group/create", gh.CreateNewGroup)
	auth.POST("/group/invite", gh.InviteUserToGroup)
	auth.POST("/group/invitation/accept", gh.AcceptInvitation)
	auth.POST("/group/invitation/reject", gh.RejectInvitation)
	auth.POST("/group/invitation/cancel", gh.CancelInvitation)
	auth.POST("/group/member/remove", gh.RemoveUserFromGroup)
	auth.POST("/group/delete", gh.DeleteGroup)

	auth.GET("/group/:id/details", gh.GetGroupDetails)
	auth.GET("/group/:id/members", gh.GetGroupMembers)
}
