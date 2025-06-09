package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kajtekajtek/forum/backend/internal/models"
)

/*
	GetUserInfo parses Gin context and returns information about the user
*/
func GetUserInfo(c *gin.Context) (models.UserInfo, error) {
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

/* 
	ParseUintParam gets server ID from URL parameters and parses it to uint
*/
func ParseUintParam(c *gin.Context, key string) (uint, error) {
	param := c.Param(key)

	param64, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse %s to uint: %w", key, err)
	}

	paramUint := uint(param64)

	return paramUint, nil
}
