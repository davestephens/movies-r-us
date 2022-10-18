package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/davestephens/movies-r-us/rest-api/database"
)


func TestPostGoodMovie(t *testing.T) {
    router := setupRouter()
	database.ConnectDatabase()
	router.POST("/movies", CreateMovie)

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
