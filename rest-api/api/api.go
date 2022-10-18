package api

import (
	"github.com/davestephens/movies-r-us/rest-api/controllers"
	"github.com/davestephens/movies-r-us/rest-api/database"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Homepage)
	r.GET("/movies", controllers.GetMovies)
	r.POST("/movies", controllers.CreateMovie)
	r.POST("/notify", controllers.PostNotification)
	r.Run()
	return r
}

func Start() {
	r := setupRouter()
	database.ConnectDatabase()
	r.Run(":8080")
}