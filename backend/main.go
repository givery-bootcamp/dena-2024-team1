package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/infrastructure/http/router"
)

func main() {
	// Setup webserver
	router := router.NewRouter()
	router.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
