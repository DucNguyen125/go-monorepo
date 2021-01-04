package mysql_util

import (
	"fmt"
	"os"

	"example/services/order-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_URI"),
		os.Getenv("MYSQL_DATABASE_NAME"))
	db, error := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if error != nil {
		return error
	}
	DB = db
	return nil
}

func AutoMigrate() error {
	if err := models.AutoMigrateTable(DB); err != nil {
		return err
	}
	return nil
}
