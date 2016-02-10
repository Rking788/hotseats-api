package db

import (
	"errors"
	"fmt"
	"github.com/rking788/hotseats-api/Godeps/_workspace/src/github.com/jinzhu/gorm"
	_ "github.com/rking788/hotseats-api/Godeps/_workspace/src/github.com/lib/pq"
	"os"
)

func GetDBConnection() (*gorm.DB, error) {

	// Retrieve required environment variables describing DB connection details
	host := os.Getenv("HOTSEATS_DB_HOST")
	name := os.Getenv("HOTSEATS_DB_NAME")
	user := os.Getenv("HOTSEATS_DB_USER")
	pass := os.Getenv("HOTSEATS_DB_PASS")
	sslMode := os.Getenv("HOTSEATS_DB_SSL_MODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	if host == "" || name == "" || user == "" || pass == "" {
		return nil, errors.New("Missing one or more DB environment variables!!")
	}

	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		host, name, user, pass, sslMode)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
