package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kajtekajtek/forum/backend/internal/database"
	"github.com/kajtekajtek/forum/backend/internal/models"
	"github.com/kajtekajtek/forum/backend/internal/utils"
)

type createMessageRequest struct {
	Content string `json:"content" binding:"required,min=1,max=500"`
}

/*
type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChannelID uint      `gorm:"not null;index" json:"channelId"`
	UserID    string    `gorm:"type:text;not null;index" json:"userId"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	// relation
	Channel   *Channel  `gorm:"foreignKey:ChannelID" json:"-"`
}
*/

func CreateMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user info & server ID from request context
		user, err := utils.GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		channelIDAny, exists := c.Get("serverID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "channel ID not found in context"})
			return
		}
		channelID := channelIDAny.(uint)

		// bind request JSON
		var req createMessageRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid payload", "details": err.Error()})
			return
		}

		// create Message record
		message := models.Message {
			ChannelID: channelID,
			UserID: user.ID,
			Content: req.Content,
		}

		if err := db.Create(&message).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create message"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": message})
	}
}

func GetMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get server ID from request's context
		channelIDAny, exists := c.Get("channelID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "channel ID not found in context"})
			return
		}
		channelID := channelIDAny.(uint)

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
