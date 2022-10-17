package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify"
	"github.com/davestephens/movies-r-us/rest-api/api"
)

func TestHomepageHandler(t *testing.T) {
    mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
    r := api.SetUpRouter()
    
	r.GET("/", Homepage)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}
