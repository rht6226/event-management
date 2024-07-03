package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDbProvider struct {
	db               *gorm.DB
	connectionConfig ConnectionConfig
	gormConfig       *gorm.Config
}

func NewPostgresDbProvider(gormConfig *gorm.Config) *PostgresDbProvider {
	provider := &PostgresDbProvider{
		connectionConfig: ConnectionConfig{},
		gormConfig:       gormConfig,
	}
	provider.connectionConfig.FromEnv()
	return provider
}

func (provider *PostgresDbProvider) GetInstance() (*gorm.DB, error) {
	if provider.db == nil {
		dsn := provider.connectionConfig.getDSN()
		var err error
		provider.db, err = provider.connectDB(dsn)
		if err != nil {
			return nil, err
		}
	}
	return provider.db, nil
}

// connectDB opens a new gorm.DB connection
func (provider *PostgresDbProvider) connectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), provider.gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}
