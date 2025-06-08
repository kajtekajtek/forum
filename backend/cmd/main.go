package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/kajtekajtek/forum/backend/internal/config"
	"github.com/kajtekajtek/forum/backend/internal/middleware"
	"github.com/kajtekajtek/forum/backend/internal/handlers"
	"github.com/kajtekajtek/forum/backend/internal/database"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := database.Initialize(config)

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.CORSOrigins,
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	router.Use(middleware.KeycloakAuthMiddleware(config))

	servers := router.Group("/api/servers")
	{
        servers.POST("", handlers.CreateServerHandler(db))
        servers.GET("", handlers.GetServerListHandler(db))
	}

	router.Run(":" + config.APIPort)
}
