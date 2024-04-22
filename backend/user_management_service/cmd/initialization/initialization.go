package initialization

import "github.com/kosuke9809/SyncMate/internal/infrastructure/database/initialize"

func Start() {
	initialize.DbInitialize()
}
