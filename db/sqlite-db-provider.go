package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dbPath string = "/tmp/event-management.db"
)

type SQLiteDbProvider struct {
	db         *gorm.DB
	gormConfig *gorm.Config
}

func NewSQLiteDbProvider(config *gorm.Config) *SQLiteDbProvider {
	return &SQLiteDbProvider{
		gormConfig: config,
	}
}

func (provider *SQLiteDbProvider) GetInstance() (*gorm.DB, error) {
	if provider.db == nil {
		db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return provider.db, nil
}
