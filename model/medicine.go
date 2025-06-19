package model

import "time"

type Medicine struct {
	ID             uint       `gorm:"primaryKey;column:id" json:"id"`
	MedicineCode   string     `gorm:"unique;column:medicine_code" json:"medicine_code"`
	Name           string     `gorm:"column:name" json:"name"`
	CategoryID     *uint      `gorm:"column:category_id" json:"category_id"`
	ManufacturerID *uint      `gorm:"column:manufacturer_id" json:"manufacturer_id"`
	Type           string     `gorm:"column:type" json:"type"`
	Description    string     `gorm:"column:description" json:"description"`
	Price          float64    `gorm:"column:price" json:"price"`
	Stock          int        `gorm:"column:stock" json:"stock"`
	Unit           string     `gorm:"column:unit" json:"unit"`
	ExpiryDate     *time.Time `gorm:"column:expiry_date" json:"expiry_date"`
	CreatedBy      string     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy      string     `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt      time.Time  `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}
