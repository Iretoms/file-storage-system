//   File Storage API:
//    version: 1.0
//    title: File Storage API
//   Schemes: http, https
//   Host: localhost:8080
//   BasePath: /api/v1
//      Consumes:
//      - application/json
//   Produces:
//   - application/json
//   swagger:meta
package main

import (
	"file-storage-system/database"
	"file-storage-system/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	loadDatabase()
	serveApp()
}

func loadDatabase() {
	database.Connect()
}

func serveApp() {
	router := gin.Default()

	publicRoutes := router.Group("/api/v1")
	routes.FileRoutes(publicRoutes)

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
