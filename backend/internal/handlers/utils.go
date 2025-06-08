package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kajtekajtek/forum/backend/internal/models"
)

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

// get server ID from URL parameters
func parseServerIDParam(c *gin.Context) (uint, error) {
	serverIDParam := c.Param("serverID")

	serverID64, err := strconv.ParseUint(serverIDParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse server id: %w", err)
	}

	serverID := uint(serverID64)

	return serverID, nil
}

