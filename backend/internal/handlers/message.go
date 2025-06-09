package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kajtekajtek/forum/backend/internal/database"
	"github.com/kajtekajtek/forum/backend/internal/models"
	"github.com/kajtekajtek/forum/backend/internal/utils"
	"github.com/kajtekajtek/forum/backend/internal/sse"
)

type createMessageRequest struct {
	Content string `json:"content" binding:"required,min=1,max=500"`
}

func CreateMessage(db *gorm.DB, manager *sse.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user info & server ID from request context
		user, err := utils.GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		// get channel ID from URL parameters and parse it
		channelID, err := utils.ParseUintParam(c, "channelID")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid channel ID"})
			return
		}

		// bind request JSON
		var req createMessageRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid payload", "details": err.Error()})
			return
		}

		// create Message record
		message := models.Message{
			ChannelID: channelID,
			UserID:    user.ID,
			Content:   req.Content,
		}

		if err := db.Create(&message).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create message"})
			return
		}

		// publish the message to channel's subscribers
		manager.Publish(channelID, message)

		c.JSON(http.StatusCreated, gin.H{"message": message})
	}
}

func GetMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get channel ID from URL parameters and parse it
		channelID, err := utils.ParseUintParam(c, "channelID")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid channel ID"})
			return
		}

		// get messages by server
		messages, err := database.QueryChannelMessages(db, channelID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to fetch messages"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messages": messages})
	}
}
