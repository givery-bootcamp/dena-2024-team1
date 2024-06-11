package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/infrastructure/database"
	"myapp/internal/infrastructure/http/router"
)

func main() {
	// Initialize database
	database.SetupDB()

	// Setup webserver
	router := router.NewRouter()
	router.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
