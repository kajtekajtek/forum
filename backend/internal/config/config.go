package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// server variables
	APIPort	string

	// database variables
	PostgresHost		string
	PostgresPort		string
	PostgresUser		string
	PostgresPassword	string
	PostgresDB			string

	// keycloak variables
	KeycloakURLs		[]string
	KeycloakRealm		string
	KeycloakClientID	string

	// CORS
	CORSOrigins			[]string
}

func Load() (*Config, error) {
	var c Config

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("warning: .env not found.")
	}

	c.APIPort = os.Getenv("API_PORT")

	c.PostgresHost 		= os.Getenv("POSTGRES_HOST")
	c.PostgresPort 		= os.Getenv("POSTGRES_PORT")
	c.PostgresUser 		= os.Getenv("POSTGRES_USER")
	c.PostgresPassword 	= os.Getenv("POSTGRES_PASSWORD")
	c.PostgresDB 		= os.Getenv("POSTGRES_DB")

	keycloakURLs      := os.Getenv("KEYCLOAK_URL")
	c.KeycloakURLs     = strings.Split(keycloakURLs, ",")
	c.KeycloakRealm	   = os.Getenv("KEYCLOAK_REALM")
	c.KeycloakClientID = os.Getenv("KEYCLOAK_CLIENT_ID")

	corsOrigins  := os.Getenv("CORS_ORIGINS")
	c.CORSOrigins = strings.Split(corsOrigins, ",")

	return &c, nil
}
