package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kajtekajtek/forum/backend/internal/config"
	"github.com/kajtekajtek/forum/backend/internal/database"
	"github.com/kajtekajtek/forum/backend/internal/models"
	"github.com/kajtekajtek/forum/backend/internal/utils"
)

func KeycloakAuth(cfg *config.Config) gin.HandlerFunc {
	// token issuer's URL (keycloak's realm adress)
	issuerURLs := make([]string, 0, len(cfg.KeycloakURLs))
	for _, url := range cfg.KeycloakURLs {
		issuerURL := fmt.Sprintf("http://%s/realms/%s",
			url,
			cfg.KeycloakRealm,
		)
		issuerURLs = append(issuerURLs, issuerURL)
	}

	// OpenID Connect provider
	provider, err := oidc.NewProvider(context.Background(), issuerURLs[0])
	if err != nil {
		panic(fmt.Sprintf("get provider: %v", err))
	}

	/* 
		token verifier initialization
		- since we are using ID Token Verifier to verify Access Tokens, 
		  keycloak's Realm should have protocol mapper configured to incldude
		  clientID in it's Access Tokens
		- we are skipping issuer check to enable multiple issuer's URLs
		- https://oauth.net/id-tokens-vs-access-tokens/
		- https://pkg.go.dev/github.com/coreos/go-oidc@v2.3.0+incompatible
	*/
	verifier := provider.Verifier(&oidc.Config{
		ClientID: cfg.KeycloakClientID, 
		SkipIssuerCheck: true,
	})

	return func(c *gin.Context) {
		// retrieve token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "no token"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid header"})
			return
		}
		rawToken := parts[1]

		// verify token
		accessToken, err := verifier.Verify(context.Background(), rawToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token"})
			return
		}

		verifiedIssuer := false
		for _, i := range issuerURLs {
			if accessToken.Issuer == i {
				verifiedIssuer = true
				break
			}
		}
		if !verifiedIssuer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token"})
			return
		}

		// read "sub" (userID) and Realm roles from claims
		var claims struct {
			Sub	        string   `json:"sub"`
			RealmAccess struct {
				Roles []string `json:"roles"`
			} `json:"realm_access"`
		}

		if err := accessToken.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "cannot parse claims"})
			return
		}

		// set user's ID and Realm roles in the Gin context
		c.Set("userID", claims.Sub)
		c.Set("userRealmRoles", claims.RealmAccess.Roles)

		c.Next()
	}
}

/*
	ServerAuth gets server ID from URL parameters and checks if user is a member of the server
*/
func ServerAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user information from request's context
		user, err := utils.GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		// get server ID from URL parameters and parse it
		serverID, err := utils.ParseUintParam(c, "serverID")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid server ID"})
			return
		}

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

		// set request's server ID in the Gin context
		c.Set("serverID", serverID)

		c.Next()
	}
}

/*
	ChannelAuth gets server ID from URL parameters and checks if user is a member of the server
*/
func ChannelAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user information from request's context
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

		// get channel's server
		var server models.Server
		err = db.Where(&models.Channel{ID: channelID}).First(&server).Error
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "channel not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "query channel error"})
		}

		// check if user is a member of the server
		isMember, err := database.IsUserMemberOfServer(db, user.ID, server.ID)
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

		// set request's server ID in the Gin context
		c.Set("serverID", server.ID)
		c.Set("channelID", channelID)

		c.Next()
	}
}
