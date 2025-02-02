package api

import (
	"fmt"
	"time"

	database "github.com/kosuke9809/SyncMate/internal/infrastructure/database/postgres"
	"github.com/kosuke9809/SyncMate/internal/interactor"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/middleware"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/router"
	"github.com/labstack/echo/v4"
)

func Start() {
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
	gh := i.NewGroupHandler()
	router.NewRouter(e, uh, gh)
	middleware.NewMiddleware(e)
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
