package models

import (
	"gorm.io/gorm"
)

func AutoMigrateTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate(&Product{}); err != nil {
		return err
	}
	return nil
}
