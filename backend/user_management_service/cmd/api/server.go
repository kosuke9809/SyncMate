package cmd

import (
	"fmt"
	"time"

	database "github.com/kosuke9809/SyncMate/internal/infrastructure/database/config"
)

func ApiServerStart() {
	fmt.Println("Starting the server")
	db, err := database.NewDBWithRetry(5, 5*time.Second)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
}
