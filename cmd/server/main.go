package main

import (
	"fmt"

	"canteen-backend/config"
	"canteen-backend/pkg/api"
	"canteen-backend/pkg/database"
	"canteen-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig("./config/config.yaml")
	database.Connect()

	router := gin.Default()
	router.Use(middleware.IPLogger())

	v1 := router.Group("/api/v1")
	{
		api.RegisterRoutes(v1)
	}

	if err := router.Run(fmt.Sprintf(":%d", config.Config.Server.Port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
