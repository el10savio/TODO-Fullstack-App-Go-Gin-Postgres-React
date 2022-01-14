module github.com/el10savio/TODO-Fullstack-App-Go-Gin-Postgres-React/backend

go 1.15

require (
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.0
	github.com/lib/pq v1.8.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/el10savio/TODO-Fullstack-App-Go-Gin-Postgres-React/backend/api => ../api
