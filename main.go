package main

import (
	"file-storage-system/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApp()
}

func loadDatabase() {
	database.Connect()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApp() {
	router := gin.Default()

	publicRoutes := router.Group("/api")
	publicRoutes.POST("/file")

	router.Run(":8080")
	fmt.Println("Server running on port 8000")
}
