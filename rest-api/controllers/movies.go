package controllers

import (
	"net/http"

	"fmt"

	"github.com/davestephens/movies-r-us/rest-api/database"
	"github.com/davestephens/movies-r-us/rest-api/models"
	"github.com/davestephens/movies-r-us/rest-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// GET /movies
// Get all movies
func GetMovies(c *gin.Context) {
	var movies []models.Movie
	//var moviesCount int64

	query := database.DB.Model(&models.Movie{})

	if c.Query("title") != "" {
		query = query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%",c.Query("title")))
	}

	if c.Query("year") != "" {
		query = query.Where("year = ?", c.Query("year"))
	}

	if c.Query("genre") != "" {
		query = query.Where("genres ILIKE ?", fmt.Sprintf("%%%s%%",c.Query("genre")))
	}

	if c.Query("actor") != "" {
		query = query.Where("actors ILIKE ?", fmt.Sprintf("%%%s%%",c.Query("actor")))
	}

	err := query.Find(&movies).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

func MovieByTitle(c *gin.Context) {
	var movies []models.Movie
	database.DB.Where("title LIKE ?", c.Param("title")).Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

func CreateMovie(c *gin.Context) {
	var movies []models.Movie

	// try and bind the input json to a slice of movieinput struct
	err := c.ShouldBindJSON(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "title"}},
		DoUpdates: clause.AssignmentColumns([]string{"year", "genres", "actors"}),
	  }).Create(&movies)
}