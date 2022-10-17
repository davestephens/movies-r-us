package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Homepage(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message":"Welcome to movies-r-us"})
}