package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() error {
	// env var setup
	viper.SetEnvPrefix("mru")
	viper.BindEnv("db_host")
	viper.BindEnv("db_user")
	viper.BindEnv("db_pass")
	viper.BindEnv("db_name")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  LogLevel: logger.Info,
		  Colorful: true,
		},
	  )

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/London",
		viper.Get("db_host"), viper.Get("db_user"), viper.Get("db_pass"), viper.Get("db_name"))

		//dsn := "root:Fishcakes448!@tcp(127.0.0.1:3306)/movie4?charset=utf8mb4&parseTime=True&loc=Local"

		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	  })

	if err != nil {
		return err
	}

	DB = database
	return nil
}