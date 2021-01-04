package models

import "time"

type Order struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id"`
	OrderCode   string    `gorm:"column:order_code"`
	OrderType   string    `gorm:"column:order_type"`
	Products    string    `gorm:"column:products"`
	OrderStatus string    `gorm:"column:order_status"`
	Quantity    int       `gorm:"column:quantity"`
	TotalPrice  int       `gorm:"column:total_price"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
