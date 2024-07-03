package db

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

func TestSQLiteDbProviderGetInstance(t *testing.T) {
	config := &gorm.Config{}
	provider := NewSQLiteDbProvider(config)

	db, err := provider.GetInstance()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if db == nil {
		t.Fatalf("expected db instance, got nil")
	}

	// Check if the database connection is open
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Clean up the database file
	defer func() {
		sqlDB.Close()
		os.Remove("/tmp/event-management.db")
	}()
}
