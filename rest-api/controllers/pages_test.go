package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
)

func TestHomepageHandler(t *testing.T) {
    expectedResponse := `{"message":"Welcome to movies-r-us"}`
    router := setupRouter()
    router.GET("/", Homepage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

func setupRouter() *gin.Engine{
    router := gin.Default()
    return router
}
