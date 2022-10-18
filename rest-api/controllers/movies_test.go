package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/davestephens/movies-r-us/rest-api/database"
	"github.com/davestephens/movies-r-us/rest-api/utils"
	"github.com/stretchr/testify/assert"
)


func TestPostGoodMovie(t *testing.T) {
    router := setupRouter()
	err := database.ConnectDatabase()
	if err != nil {
		utils.Logger.Panic("Can't connect to test database :(")
	}
	router.POST("/movies", PostMovie)

	// open test data
    json, err := os.Open("../testdata/goodmovie.json")
	if err != nil {
		panic(err)
	}

	// post
	req, _ := http.NewRequest("POST", "/movies", json)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
