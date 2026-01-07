package main

import (
	"log"

	"github.com/Shubham7985/user-module-api/config"
	"github.com/Shubham7985/user-module-api/middlewares"
	"github.com/Shubham7985/user-module-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Add logger middleware
	r.Use(middlewares.Logger())

	// Routes

	// Connect to Supabase Postgres
	config.ConnectDB()

	// Setup router
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080") // port 8080
}
