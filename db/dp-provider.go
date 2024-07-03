package db

import "gorm.io/gorm"

type DBProvider interface {
	GetInstance() (*gorm.DB, error)
}
