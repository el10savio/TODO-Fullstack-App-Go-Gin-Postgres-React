package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ListItem struct {
	Id   string `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

var db *sql.DB
var err error

func SetupPostgres() {
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/todo?sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	log.Println("connected to postgres")
}

// CRUD: Create Read Update Delete API Format
// Add invalid output Gin responses 404, 403, etc

// List all todo items
func TodoItems(c *gin.Context) {
	// Use SELECT Query to obtain all rows
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	defer rows.Close()

	// Get all rows and add into items
	items := make([]ListItem, 0)
	for rows.Next() {
		// Individual row processing
		item := ListItem{}
		if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	// Return JSON objet of all rows
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// Create todo item
func CreateTodoItem(c *gin.Context) {
	item := c.Param("item")

	if len(item) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an item"})
	} else {
		var TodoItem ListItem

		TodoItem.Item = item
		TodoItem.Done = false

		_, err := db.Query("INSERT INTO list(item, done) VALUES($1, $2);", TodoItem.Item, TodoItem.Done)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully create todo item"})
	}
}

// Update todo item
func UpdateTodoItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update todo item"})
}

// Delete todo item
func DeleteTodoItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "delete todo item"})
}
