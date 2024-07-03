package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetInstance(t *testing.T) {
	// Mock configuration
	envVars := map[string]string{
		hostName: "localhost",
		dbName:   "testdb",
		password: "password",
		user:     "user",
		port:     "5432",
		sslMode:  "disable",
		timeZone: "UTC",
	}
	for key, value := range envVars {
		os.Setenv(key, value)
		defer os.Unsetenv(key)
	}
	mockGormConfig := &gorm.Config{}
	provider := NewPostgresDbProvider(mockGormConfig)
	provider.db = &gorm.DB{
		Config: mockGormConfig,
	}
	db, err := provider.GetInstance()
	assert.NoError(t, err, "Expected no error when initializing database connection")
	assert.NotNil(t, db, "Expected db instance to be initialized")
}
