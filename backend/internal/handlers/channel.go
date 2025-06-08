package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kajtekajtek/forum/backend/internal/database"
	"github.com/kajtekajtek/forum/backend/internal/models"
	"github.com/kajtekajtek/forum/backend/internal/utils"
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
				"error": "server ID not fount in context"})
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
		// get user information and server ID from request's context
		user, err := utils.GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		serverIDAny, exists := c.Get("serverID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "server ID not fount in context"})
			return
		}
		serverID := serverIDAny.(uint)

		// check if user is a member of the server
		isMember, err := database.IsUserMemberOfServer(db, user.ID, serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "membership check failed"})
			return
		}
		if !isMember {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "not a member of this server"})
			return
		}
		
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
