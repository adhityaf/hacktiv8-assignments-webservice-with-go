package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	// "gorm.io/gorm"
)

const (
	DB_HOST  = "localhost"
	DB_PORT  = "5432"
	DB_USER  = "pgadmin"
	DB_PASS  = "pgadmin"
	DB_NAME  = "db-assign2"
	APP_PORT = ":8080"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	log.Default().Println("DB connected successfully")
	db.AutoMigrate(models.Item{}, models.Order{})
	return db, nil
}
