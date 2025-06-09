package handlers

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/kajtekajtek/forum/backend/internal/models"
	"github.com/kajtekajtek/forum/backend/internal/database"
	"github.com/kajtekajtek/forum/backend/internal/utils"
	"gorm.io/gorm"
)

type createServerRequest struct {
	Name string `json:"name" binding:"required,min=1,max=30"`
}

func CreateServer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.GetUserInfo(c)
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

func GetServerList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		isAdmin := slices.Contains(user.RealmRoles, "admin")
		isMod := slices.Contains(user.RealmRoles, "moderator")
		
		var servers []models.Server

		// if user has role admin or moderator, return all servers
		if isAdmin || isMod {
			servers, err = database.QueryAllServers(db)	
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "failed to query servers"})
                return
			}
		// else, query servers by user ID
		} else {
			servers, err = database.QueryUserServers(db, user.ID)	
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "failed to query user servers"})
				return
			}
		}

		// return JSON
		c.JSON(http.StatusOK, gin.H{
			"servers": servers,
		})
	}
}
