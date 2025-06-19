package repository

import (
	"toko_obat/model"
	"toko_obat/repo/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) *MedicineRepository {
	return &MedicineRepository{db: db}
}
func (m *MedicineRepository) GetAllMedicine(filter request.GetFilter) ([]model.Medicine, int64, int64, error) {
	var model []model.Medicine
	var recordtotal int64
	var recordtotalfiltered int64
	var sort bool // variable for sort desc or asc

	q := m.db.Table(`medicines`).
		Joins(`left join categories on categories.id = medicines.category_id `).
		Joins(`left join manufacturers on manufacturers.id = medicines.manufacturer_id`).
		Count(&recordtotal) // count before filter
	// couunt after filter
	q.Count(&recordtotalfiltered)

	// order by
	if filter.Dir == "desc" || filter.Dir == "" {
		sort = true
	} else {
		sort = false
	}

	q.Offset(filter.PageOffset).
		Limit(filter.PageLimit).
		Order(clause.OrderByColumn{Column: clause.Column{Name: MapMedicine(filter.Field)}, Desc: sort}).
		Select(`*`).
		Find(&model)

	return model, recordtotal, recordtotalfiltered, nil
}

// map for sort column list
func MapMedicine(value string) string {
	var mp map[string]string = map[string]string{
		"id":              "medicines.id",
		"medicine_code":   "medicine_code",
		"name":            "name",
		"category_id":     "category_id",
		"manufacturer_id": "manufacturer_id",
		"type":            "type",
		"description":     "description",
		"price":           "price",
		"stock":           "stock",
		"unit":            "unit",
		"expiry_date":     "expiry_date",
		"created_by":      "created_by",
		"updated_by":      "updated_by",
	}

	return mp[value]
}

func (m *MedicineRepository) GetMedicineByID(id uint) (*model.Medicine, error) {
	var medicine model.Medicine
	err := m.db.Table("medicines").
		Joins("left join categories on categories.id = medicines.category_id").
		Joins("left join manufacturers on manufacturers.id = medicines.manufacturer_id").
		Where("medicines.id = ?", id).
		First(&medicine).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}
