package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/coreos/go-oidc"

	"github.com/kajtekajtek/forum/backend/internal/config"
)

func KeycloakAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	// token issuer's URL (keycloak's realm adress)
	issuer := fmt.Sprintf("http://%s:%s/realms/%s", 
		cfg.KeycloakHost, 
		cfg.KeycloakPort, 
		cfg.KeycloakRealm,
	)

	// OpenID Connect provider
	provider, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		panic(fmt.Sprintf("get provider: %v", err))
	}

	/* 
		token verifier initialization
		- since we are using ID Token Verifier to verify Access Tokens, 
		  keycloak's Realm should have protocol mapper configured to incldude
		  clientID in it's Access Tokens
		- https://oauth.net/id-tokens-vs-access-tokens/
		- https://pkg.go.dev/github.com/coreos/go-oidc@v2.3.0+incompatible
	*/
	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.KeycloakClientID})

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
		c.Set("roles", claims.RealmAccess.Roles)

		c.Next()
	}
}
