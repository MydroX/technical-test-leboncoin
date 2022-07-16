package main

import (
	"log"
	"os"

	"github.com/MydroX/leboncoin-technical-test/src/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	r.POST("/", handlers.Home)

	r.Run(":3000")
}
