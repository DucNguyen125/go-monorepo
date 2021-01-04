package models

import "time"

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement;column:id"`
	FirstName  string    `gorm:"column:first_name"`
	LastName   string    `gorm:"column:last_name"`
	Email      string    `gorm:"column:email;unique"`
	Password   string    `gorm:"column:password"`
	FacebookId string    `gorm:"column:facebook_id"`
	GoogleId   string    `gorm:"column:google_id"`
	Avatar     string    `gorm:"column:avatar"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
