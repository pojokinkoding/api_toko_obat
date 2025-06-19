package repository

import (
	"toko_obat/model"
	"toko_obat/repo/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ManufacturerRepository struct {
	db *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{db: db}
}

func (r *ManufacturerRepository) GetAllManufacturer(filter request.GetFilter) ([]model.Manufacturer, int64, int64, error) {
	var manufacturers []model.Manufacturer
	var recordtotal int64
	var recordtotalfiltered int64
	var sort bool

	q := r.db.Table("manufacturers").Count(&recordtotal)
	q.Count(&recordtotalfiltered)

	if filter.Dir == "desc" || filter.Dir == "" {
		sort = true
	} else {
		sort = false
	}

	q.Offset(filter.PageOffset).
		Limit(filter.PageLimit).
		Order(clause.OrderByColumn{Column: clause.Column{Name: MapManufacturer(filter.Field)}, Desc: sort}).
		Select("*").
		Find(&manufacturers)

	return manufacturers, recordtotal, recordtotalfiltered, nil
}

func MapManufacturer(value string) string {
	var mp = map[string]string{
		"id":         "manufacturers.id",
		"name":       "name",
		"address":    "address",
		"contact":    "contact",
		"created_by": "created_by",
		"updated_by": "updated_by",
		"created_at": "created_at",
		"updated_at": "updated_at",
	}
	return mp[value]
}

func (r *ManufacturerRepository) GetManufacturerByID(id uint) (*model.Manufacturer, error) {
	var manufacturer model.Manufacturer
	err := r.db.Table("manufacturers").Where("manufacturers.id = ?", id).First(&manufacturer).Error
	if err != nil {
		return nil, err
	}
	return &manufacturer, nil
}
