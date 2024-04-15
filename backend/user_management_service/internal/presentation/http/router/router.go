package router

import (
	"github.com/kosuke9809/SyncMate/internal/presentation/http/handler"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/middleware"
	"github.com/labstack/echo/v4"
)

func NewRouter(h handler.IndexHandler) *echo.Echo {
	e := echo.New()
	e = middleware.NewMiddleware(e)

	return e
}
