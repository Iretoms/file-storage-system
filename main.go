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

	publicRoutes := router.Group("/api")
	routes.FileRoutes(publicRoutes)

	router.Run(":8080")
	fmt.Println("Server running on port 8000")
}
