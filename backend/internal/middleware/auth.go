package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"

	"github.com/kajtekajtek/forum/backend/internal/config"
)

func KeycloakAuthMiddleware(c *config.Config) gin.HandlerFunc {
	issuer := fmt.Sprintf("%s/realms/%s", c.KeycloakURL, c.KeycloakRealm)

	provider, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		panic(fmt.Sprintf("get provider: %v", err))
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: c.KeycloakClientID})

	return func(c *gin.Context) {
		// retrieve and verify token
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
		token := parts[1]

		idToken, err := verifier.Verify(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token"})
			return
		}

		// retrieve user ID from claims
		var claims struct {
			Sub	string `json:"sub"`
		}
		if err := idToken.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "cannot parse claims"})
			return
		}

		c.Set("userID", claims.Sub)
		c.Next()
	}
}
