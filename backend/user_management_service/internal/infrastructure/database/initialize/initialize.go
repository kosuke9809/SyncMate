package initialize

import (
	"fmt"
	"os"
	"time"

	database "github.com/kosuke9809/SyncMate/internal/infrastructure/database/postgres"
)

func DbInitialize() {
	dbConn, err := database.NewDBWithRetry(5, 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer fmt.Println("Suceessfully initialized the database")
	defer database.CloseDB(dbConn)

	var count int64
	if err := dbConn.Table("roles").Count(&count).Error; err != nil {
		panic("failed to check roles table: " + err.Error())
	}

	if count == 0 {
		sqlBytes, err := os.ReadFile("./internal/infrastructure/database/initialize/sql/init.sql")
		if err != nil {
			panic("failed to read SQL file: " + err.Error())
		}
		initSQL := string(sqlBytes)

		if err := dbConn.Exec(initSQL).Error; err != nil {
			panic("failed to execute SQL: " + err.Error())
		}
	} else {
		fmt.Println("Database is already initialized")
	}
}
