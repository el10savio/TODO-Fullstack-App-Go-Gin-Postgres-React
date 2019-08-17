package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func displayTable() {
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	// Get all rows and add into items
	items := make([]ListItem, 0)
	for rows.Next() {
		// Individual row processing
		item := ListItem{}
		if err := rows.Scan(&item.Id, &item.Item, &item.Done); err != nil {
			fmt.Println(err.Error())
		}
		items = append(items, item)
	}

	fmt.Println("items:", items)
}

func emptyTable() {
	db.Exec("DELETE from list;")
	db.Exec("ALTER SEQUENCE list_id_seq RESTART WITH 1;")
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

// Test for successfull create
// response from /item/create
func TestItemCreate(t *testing.T) {
	emptyTable()

	// Expected body
	body := gin.H{
		"items": ListItem{
			Id:   "",
			Item: "Test-API",
			Done: false,
		},
	}

	w := performRequest(router, "GET", "/item/create/Test-API")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]ListItem
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["items"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["items"], value)
}

// Test for successfull creates
// response for multiple items
// from /item/create
func TestItemsCreate(t *testing.T) {
	emptyTable()

	// Expected body
	body := gin.H{
		"items": []ListItem{
			{
				Id:   "1",
				Item: "Test-API",
				Done: false,
			},
			{
				Id:   "2",
				Item: "Test-DB",
				Done: false,
			},
		},
	}

	w1 := performRequest(router, "GET", "/item/create/Test-API")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := performRequest(router, "GET", "/item/create/Test-DB")
	assert.Equal(t, http.StatusOK, w2.Code)

	w3 := performRequest(router, "GET", "/items")
	assert.Equal(t, http.StatusOK, w3.Code)

	var response map[string][]ListItem
	err := json.Unmarshal([]byte(w3.Body.String()), &response)
	value, exists := response["items"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["items"], value)
}

// Test for successfull delete
// from /item/delete
func TestItemDelete(t *testing.T) {
	emptyTable()

	// Expected body
	body := gin.H{
		"items": []ListItem{
			{
				Id:   "2",
				Item: "Test-DB",
				Done: false,
			},
		},
	}

	w1 := performRequest(router, "GET", "/item/create/Test-API")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := performRequest(router, "GET", "/item/create/Test-DB")
	assert.Equal(t, http.StatusOK, w2.Code)

	w3 := performRequest(router, "GET", "/item/delete/1")
	assert.Equal(t, http.StatusOK, w3.Code)

	w4 := performRequest(router, "GET", "/items")
	assert.Equal(t, http.StatusOK, w4.Code)

	var response map[string][]ListItem
	err := json.Unmarshal([]byte(w4.Body.String()), &response)
	value, exists := response["items"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["items"], value)
}

// Test for successfull update
// from /item/update
func TestItemUpdate(t *testing.T) {
	emptyTable()

	// Expected body
	body := gin.H{
		"items": []ListItem{
			{
				Id:   "1",
				Item: "Test-API",
				Done: true,
			},
		},
	}

	w1 := performRequest(router, "GET", "/item/create/Test-API")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := performRequest(router, "GET", "/item/update/1/true")
	assert.Equal(t, http.StatusOK, w2.Code)

	w3 := performRequest(router, "GET", "/items")
	assert.Equal(t, http.StatusOK, w3.Code)

	var response map[string][]ListItem
	err := json.Unmarshal([]byte(w3.Body.String()), &response)
	value, exists := response["items"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["items"], value)
}
