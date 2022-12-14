package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/davestephens/movies-r-us/rest-api/models"
	"github.com/davestephens/movies-r-us/rest-api/s3"
	"github.com/davestephens/movies-r-us/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func PostNotification(c *gin.Context) {
	var notification models.Notification

	// try and bind the input json to a slice of Notification struct
	err := c.ShouldBindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.Logger.Infof("Received new file notification %s", notification)

	jsonFile := s3.DownloadFile(notification)

	utils.Logger.Infof("Opening %s for reading", jsonFile.Name())
	file, err := os.Open(jsonFile.Name())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	byteValue, _ := ioutil.ReadAll(file)

	var movies []models.Movie

	utils.Logger.Info("Attempting to unmarshall json")
	// try and bind the input json to a slice of movieinput struct
	json.Unmarshal(byteValue, &movies)

	s := string(byteValue)
	utils.Logger.Info("file contents %s", s)

	_, err = CreateMovie(movies); if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	  c.JSON(http.StatusOK, gin.H{"message":"OK"})
}