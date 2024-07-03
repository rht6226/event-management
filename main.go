package main

import (
	"fmt"
	"time"

	"os"

	"github.com/rht6226/event-management-app/db"
	"gorm.io/gorm"
)

const (
	deploymentEnvironment string = "DEPLOYMENT_ENVIRONMENT"
)

var (
	database *gorm.DB
)

func init() {
	var provider db.DBProvider

	// load env var and select provider
	if envType := os.Getenv(deploymentEnvironment); envType == "production" {
		provider = db.NewPostgresDbProvider(&gorm.Config{})
	} else {
		provider = db.NewSQLiteDbProvider(&gorm.Config{})
	}

	database, err := provider.GetInstance()
	if err != nil {
		panic(err)
	}
	database.AutoMigrate()
}

func main() {
	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db...")
	fmt.Println("Application started...")
	time.Sleep(10 * time.Hour)
}
