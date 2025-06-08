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
	if err != nil {
		log.Fatalf("initialize database: %v", err)
	}

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.CORSOrigins,
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	router.Use(middleware.KeycloakAuth(config))

	servers := router.Group("/api/servers")
	{
		// api/servers
        servers.POST("", handlers.CreateServer(db))
        servers.GET("", handlers.GetServerList(db))

		server := router.Group("/:serverID", middleware.ServerAuth(db))
		{
			// api/servers/:serverID
			server.POST("/channels", handlers.CreateChannel(db))
			server.GET("/channels", handlers.GetChannelList(db))
		}
	}

	router.Run(":" + config.APIPort)
}
