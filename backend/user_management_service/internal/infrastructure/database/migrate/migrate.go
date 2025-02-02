package migrate

import (
	"fmt"
	"time"

	"github.com/kosuke9809/SyncMate/internal/domain/model"
	database "github.com/kosuke9809/SyncMate/internal/infrastructure/database/postgres"
)

func AutoMigrate() {
	dbConn, err := database.NewDBWithRetry(5, 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer fmt.Println("Suceessfully migrated the database")
	defer database.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&model.User{}, &model.Group{}, &model.UserGroup{}, &model.Profile{}, &model.Role{}, &model.Permission{}, &model.Invitation{}); err != nil {
		panic(err)
	}
}
