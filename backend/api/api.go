package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CRUD: Create Read Update Delete API Format

// List all todo items
func TodoItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list todo items"})
}

// Create todo item
func CreateTodoItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create todo item"})
}

// Update todo item
func UpdateTodoItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update todo item"})
}

// Delete todo item
func DeleteTodoItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "delete todo item"})
}
