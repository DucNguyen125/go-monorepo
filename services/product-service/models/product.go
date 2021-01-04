package models

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id"`
	ProductCode string    `gorm:"column:product_code"`
	ProductName string    `gorm:"column:product_name"`
	Price       int       `gorm:"column:price"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
