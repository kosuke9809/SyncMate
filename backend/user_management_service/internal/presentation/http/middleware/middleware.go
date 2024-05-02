package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewMiddleware(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"localhost:3000", os.Getenv("FRONTEND_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	// e.Use(echojwt.WithConfig(
	// 	echojwt.Config{
	// 		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
	// 		TokenLookup: "header: Authorization",
	// 		ContextKey:  "claims",
	// 	},
	// ))
	return e
}
