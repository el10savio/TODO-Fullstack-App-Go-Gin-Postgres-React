package main

import (
	"net/http"

	api "./api"

	"github.com/gin-gonic/gin"
)

// Function called for index
func indexView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO APP"})
}

// Setup Gin Routes
func SetupRoutes() {
	// Use Gin as router
	router := gin.Default()

	// Set route for index
	router.GET("/", indexView)

	// Set routes for API
	// Update to POST, UPDATE, DELETE etc
	router.GET("/items", api.TodoItems)
	router.GET("/item/create/:item", api.CreateTodoItem)
	router.GET("/item/update/:id/:done", api.UpdateTodoItem)
	router.DELETE("/item/delete/:id", api.DeleteTodoItem)

	// Set up Gin Server
	router.Run(":8081")
}

// Main function
func main() {
	api.SetupPostgres()
	SetupRoutes()
}
