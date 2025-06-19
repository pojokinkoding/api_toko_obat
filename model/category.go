package model

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}
