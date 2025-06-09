package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kajtekajtek/forum/backend/internal/database"
	"github.com/kajtekajtek/forum/backend/internal/models"
)

type createChannelRequest struct {
	Name string `json:"name" binding:"required,min=1,max=30"`
}

func CreateChannel(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get server's ID from request's context
		serverIDAny, exists := c.Get("serverID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "server ID not found in context"})
			return
		}
		serverID := serverIDAny.(uint)

		// bind request JSON
		var req createChannelRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid payload", "details": err.Error()})
			return
		}

		// create Channel record
		channel := models.Channel {
			Name: req.Name,
			ServerID: serverID,
		}

		if err := db.Create(&channel).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create channel"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"channel": channel})
	}
}

func GetChannelList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get server ID from request's context
		serverIDAny, exists := c.Get("serverID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "server ID not found in context"})
			return
		}
		serverID := serverIDAny.(uint)

		// get channels by server
		channels, err := database.QueryServerChannels(db, serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to fetch channels"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"channels": channels})
	}
}
