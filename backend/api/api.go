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

// List all todo items
func TodoItems(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	defer rows.Close()

	items := make([]ListItem, 0)
	for rows.Next() {
		item := ListItem{}
		if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
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
