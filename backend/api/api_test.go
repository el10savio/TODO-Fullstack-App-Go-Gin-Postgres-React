package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func emptyTable() {
	db.Exec("DELETE from list")
}

// Setup Gin Routes
func SetupRoutes() *gin.Engine {
	// Use Gin as router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Set routes for API
	router.GET("/items", TodoItems)
	router.GET("/item/create/:item", CreateTodoItem)
	router.GET("/item/update/:id/:done", UpdateTodoItem)
	router.GET("/item/delete/:id", DeleteTodoItem)

	// Set up Gin Server
	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMain(t *testing.T) {
	SetupPostgres()
	router = SetupRoutes()
}

// Test for successfull GET
// response from /items
// with no elements
func TestItemsGet(t *testing.T) {
	emptyTable()

	// Expected body
	body := gin.H{
		"items": []ListItem{},
	}

	w := performRequest(router, "GET", "/items")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]ListItem
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["items"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["items"], value)
}
