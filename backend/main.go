package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO APP"})
}

func setupRoutes() {
	router := gin.Default()
	router.GET("/", indexView)
	router.Run(":8081")
}

func main() {
	setupRoutes()
}
