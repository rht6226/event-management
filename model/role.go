package model

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
)

type Role string

const (
	UserRole        Role = "USER"
	AdminRole       Role = "ADMIN"
	CoordinatorRole Role = "COORDINATOR"
)

// Scanner Interface for Role Enum
func (r *Role) Scan(value interface{}) error {
	*r = Role(value.(string))
	return nil
}

func (r Role) Value() (driver.Value, error) {
	switch r {
	case UserRole, AdminRole, CoordinatorRole:
		return string(r), nil
	}
	return nil, fmt.Errorf("invalid role: %s", r)
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	return createRoleEnum(tx)
}

func createRoleEnum(tx *gorm.DB) error {
	return tx.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN CREATE TYPE role AS ENUM ('USER', 'ADMIN', 'COORDINATOR'); END IF; END $$;").Error
}
