package handlers

import (
	"net/http"
	"fmt"


	"github.com/gin-gonic/gin"
	"github.com/kajtekajtek/forum/backend/internal/models"
	"gorm.io/gorm"
)

type createServerRequest struct {
	Name string `json:"name" binding:"required,min=1"`
}

func CreateServerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		// bind request JSON
		var req createServerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid payload", "details": err.Error()})
			return
		}

		// create Server record
		server := models.Server{
			Name: req.Name,
		}

		if err := db.Create(&server).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create server"})
			return
		}

		// create Membership record
		membership := models.Membership{
			UserID: user.ID,
			ServerID: server.ID,
			Role: "admin",	// assign server admin role to server's creator
		}

		if err := db.Create(&membership).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create membership"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"server": server,
			"membership": membership,
		})
	}
}

func GetServerListHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}
		
		// query user's memberships from database
		var memberships []models.Membership
		err = db.Where("user_id = ?", user.ID).Find(&memberships).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to query user's memberships"})
			return
		}

		// get server IDs from memberships
		serverIDs := make([]uint, 0, len(memberships))
		for _, m := range memberships {
			serverIDs = append(serverIDs, m.ServerID)
		}

		// query servers with given IDs from database
		var servers []models.Server
		if len(serverIDs) > 0 {
			err := db.Where("id IN ?", serverIDs).Find(&servers).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "failed to query servers"})
				return
			}
		}

		// return JSON
		c.JSON(http.StatusOK, gin.H{
			"servers": servers,
		})
	}
}

/*
	getUserInfo parses Gin context and returns information about user
*/
func getUserInfo(c *gin.Context) (models.UserInfo, error) {
	var user models.UserInfo

	// get user ID
	userID, exists := c.Get("userID")
	if !exists {
		return models.UserInfo{}, fmt.Errorf("no user ID in context")
	}
	user.ID = userID.(string)

	// get user's Realm roles
	userRealmRoles, exists := c.Get("userRealmRoles")
	if !exists {
		return models.UserInfo{}, fmt.Errorf("no user Realm roles in context")
	}
	user.RealmRoles = userRealmRoles.([]string)

	return user, nil
}
