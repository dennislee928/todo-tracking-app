package database

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection to PostgreSQL (Supabase compatible).
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DB returns the underlying *sql.DB for migrations if needed.
func DB(gormDB *gorm.DB) (*sql.DB, error) {
	return gormDB.DB()
}
