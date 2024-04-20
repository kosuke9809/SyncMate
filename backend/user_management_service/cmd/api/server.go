package api

import (
	"fmt"
	"time"

	database "github.com/kosuke9809/SyncMate/internal/infrastructure/database/config"
	"github.com/kosuke9809/SyncMate/internal/interactor"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/middleware"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/router"
	"github.com/labstack/echo/v4"
)

func ApiServerStart() {
	fmt.Println("Starting the server")
	db, err := database.NewDBWithRetry(5, 5*time.Second)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer database.CloseDB(db)
	e := echo.New()
	i := interactor.NewInteractor(db)
	uh := i.NewUserHandler()
	router.NewRouter(e, uh)
	middleware.NewMiddleware(e)
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
