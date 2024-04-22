package migration

import "github.com/kosuke9809/SyncMate/internal/infrastructure/database/migrate"

func Start() {
	migrate.AutoMigrate()
}
