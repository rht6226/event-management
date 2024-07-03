package db

import (
	"os"
	"testing"
)

func TestConnectionConfig(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected ConnectionConfig
	}{
		{
			name: "Full Config",
			envVars: map[string]string{
				hostName: "localhost",
				dbName:   "testdb",
				password: "password",
				user:     "user",
				port:     "5432",
				sslMode:  "disable",
				timeZone: "UTC",
			},
			expected: ConnectionConfig{
				Host:         "localhost",
				DatabaseName: "testdb",
				Password:     "password",
				User:         "user",
				Port:         "5432",
				SSLMode:      "disable",
				TimeZone:     "UTC",
			},
		},
		{
			name: "Partial Config",
			envVars: map[string]string{
				hostName: "localhost",
				dbName:   "testdb",
			},
			expected: ConnectionConfig{
				Host:         "localhost",
				DatabaseName: "testdb",
				Password:     "",
				User:         "",
				Port:         "",
				SSLMode:      "",
				TimeZone:     "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}
			config := ConnectionConfig{}
			config.FromEnv()

			if config.Host != tt.expected.Host {
				t.Errorf("Expected Host to be %s, got %s", tt.expected.Host, config.Host)
			}
			if config.DatabaseName != tt.expected.DatabaseName {
				t.Errorf("Expected DatabaseName to be %s, got %s", tt.expected.DatabaseName, config.DatabaseName)
			}
			if config.Password != tt.expected.Password {
				t.Errorf("Expected Password to be %s, got %s", tt.expected.Password, config.Password)
			}
			if config.User != tt.expected.User {
				t.Errorf("Expected User to be %s, got %s", tt.expected.User, config.User)
			}
			if config.Port != tt.expected.Port {
				t.Errorf("Expected Port to be %s, got %s", tt.expected.Port, config.Port)
			}
			if config.SSLMode != tt.expected.SSLMode {
				t.Errorf("Expected SSLMode to be %s, got %s", tt.expected.SSLMode, config.SSLMode)
			}
			if config.TimeZone != tt.expected.TimeZone {
				t.Errorf("Expected TimeZone to be %s, got %s", tt.expected.TimeZone, config.TimeZone)
			}
		})
	}
}

func TestConnectionConfigGetDSN(t *testing.T) {
	config := ConnectionConfig{
		Host:         "localhost",
		DatabaseName: "testdb",
		Password:     "password",
		User:         "user",
		Port:         "5432",
		SSLMode:      "disable",
		TimeZone:     "UTC",
	}

	expectedDSN := "host=localhost user=user password=password dbname=testdb port=5432 sslmode=disable TimeZone=UTC"

	dsn := config.getDSN()

	if dsn != expectedDSN {
		t.Errorf("Expected DSN to be '%s', got '%s'", expectedDSN, dsn)
	}
}
