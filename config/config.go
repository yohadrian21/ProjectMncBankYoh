// config/config.go
package config

import (
	"gorm.io/driver/mysql"    // For MySQL
	"gorm.io/driver/postgres" // For PostgreSQL
	"gorm.io/gorm"
	// For MongoDB
)

func ConnectPostgres(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func ConnectMySQL(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
