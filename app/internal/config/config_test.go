package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	// Set environment variables
	os.Setenv("PSQL_USERNAME", "testuser")
	os.Setenv("PSQL_PASSWORD", "testpass")
	os.Setenv("PSQL_DATABASE", "testdb")
	os.Setenv("JWT_SECRET_KEY", "secret")
	os.Setenv("JWT_TOKEN_DURATION", "1h")

	// Reset instance to nil to allow reinitialization
	// This is optional and a bit of a hack to bypass the singleton pattern
	instance = nil

	// Get configuration
	config := GetConfig()

	// Assertions
	assert.Equal(t, "testuser", config.PostgreSQL.PostgreUsername)
	assert.Equal(t, "testpass", config.PostgreSQL.Password)
	assert.Equal(t, "testdb", config.PostgreSQL.Database)
	assert.Equal(t, "secret", config.Jwt.SecretKey)
	assert.Equal(t, "1h", config.Jwt.TokenDuration)

	// Cleanup: Unset environment variables after the test
	os.Unsetenv("PSQL_USERNAME")
	os.Unsetenv("PSQL_PASSWORD")
	os.Unsetenv("PSQL_DATABASE")
	os.Unsetenv("JWT_SECRET_KEY")
	os.Unsetenv("JWT_TOKEN_DURATION")
}
