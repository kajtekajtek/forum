package main

import (
	"log"

	"github.com/gin-gonic/gin"

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

	router.Use(middleware.KeycloakAuthMiddleware(config))

	servers := router.Group("/api/servers")
	{
        servers.POST("", handlers.CreateServerHandler(db))
        servers.GET("", handlers.GetServerListHandler(db))
	}
}
