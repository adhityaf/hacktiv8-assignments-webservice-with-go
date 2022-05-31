package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DSN for Postgresql
// const (
// 	DB_HOST  = "localhost"
// 	DB_PORT  = "5432"
// 	DB_USER  = "pgadmin"
// 	DB_PASS  = "pgadmin"
// 	DB_NAME  = "db-assign2"
// 	APP_PORT = ":8080"
// )

// DSN for MySQL
const (
	DB_HOST  = "localhost"
	DB_PORT  = "3306"
	DB_USER  = "root"
	DB_PASS  = ""
	DB_NAME  = "db-assign2"
	TIMEOUT  = "10s"
	APP_PORT = ":8888"
)

func ConnectDB(driver string) (db *gorm.DB, err error) {
	if driver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME, TIMEOUT)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("fail to connect to database, error=" + err.Error())
		}
	} else if driver == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("driver %s not supported", driver)
	}

	log.Default().Println("DB connected successfully")
	db.AutoMigrate(models.Order{}, models.Item{})
	return db, nil
}
