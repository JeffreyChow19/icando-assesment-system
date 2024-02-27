package lib

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config *Config) (*Database, error) {
	host := config.DatabaseHost
	port := config.DatabasePort
	username := config.DatabaseUsername
	password := config.DatabasePassword
	dbName := config.DatabaseName
	dsn := getDatabaseDsn(host, port, username, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to DB")
	}
	return &Database{
		DB: db,
	}, nil
}

func NewTestDatabase(config *Config) (*Database, error) {
	host := config.TestDatabaseHost
	port := config.TestDatabasePort
	username := config.TestDatabaseUsername
	password := config.TestDatabasePassword
	dbName := config.TestDatabaseName
	dsn := getDatabaseDsn(host, port, username, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to Test DB")
	}
	return &Database{
		DB: db,
	}, nil
}

func getDatabaseDsn(host, port, username, password, dbName string) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username,
		password, dbName, port,
	)
}
