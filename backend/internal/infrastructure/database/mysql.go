package database

import (
	"context"
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/infrastructure/database/ent"
	"os"
)

// Database Setup
// !!! You have to call this function after config setup
func SetupDB() *ent.Client {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	user := config.DBUsername
	password := config.DBPassword
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client.Debug()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
