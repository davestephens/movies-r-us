package main

import (
	"github.com/davestephens/movies-r-us/rest-api/api"
	"github.com/davestephens/movies-r-us/rest-api/database"
	"github.com/davestephens/movies-r-us/rest-api/models"
	"github.com/davestephens/movies-r-us/rest-api/utils"
)

func main() {
	utils.InitLogger()
	utils.Logger.Sync()

	if err := database.ConnectDatabase(); err != nil {
		utils.Logger.Panic("Can't connect to the database. Bad times :(")
	}

	database.DB.AutoMigrate(&models.Movie{} )

	utils.Logger.Info("Starting REST API server")
	api.Start()
}