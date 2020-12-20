package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DB Database to export and used throughout the project
var DB *gorm.DB

func Setup() {
	databaseConfig := viper.Get("database").(map[string]interface{})
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s DB.name=%s sslmode=disable", databaseConfig["host"], databaseConfig["port"], databaseConfig["user"], databaseConfig["pass"], databaseConfig["schema"])
	var err error
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}
	for _, model := range models {

		DB.AutoMigrate(model)
		DB.Debug()
	}
}
