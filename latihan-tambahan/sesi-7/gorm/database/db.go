package database

import (
	"fmt"
	"log"
	"sesi-7-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "pgadmin"
	DB_PASSWORD = "pgadmin"
	DB_NAME     = "db-go-psql"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Default().Println("connection db success")

	err = migration(db)
	if err != nil {
		panic(err)
	}

	return db
}

func migration(db *gorm.DB) error {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(models.Product{})
	if err != nil {
		return err
	}
	return nil
}
