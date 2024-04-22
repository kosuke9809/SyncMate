package main

import (
	"github.com/kosuke9809/SyncMate/cmd/api"
	"github.com/kosuke9809/SyncMate/cmd/initialization"

	"github.com/kosuke9809/SyncMate/cmd/migration"
)

func main() {
	migration.Start()
	initialization.Start()
	api.Start()
}
