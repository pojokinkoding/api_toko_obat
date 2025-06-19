package repository

import (
	"toko_obat/model"
	"toko_obat/repo/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategory(filter request.GetFilter) ([]model.Category, int64, int64, error) {
	var model []model.Category
	var recordtotal int64
	var recordtotalfiltered int64
	var sort bool

	q := r.db.Table("categories").Count(&recordtotal)
	q.Count(&recordtotalfiltered)

	if filter.Dir == "desc" || filter.Dir == "" {
		sort = true
	} else {
		sort = false
	}

	q.Offset(filter.PageOffset).
		Limit(filter.PageLimit).
		Order(clause.OrderByColumn{Column: clause.Column{Name: MapCategory(filter.Field)}, Desc: sort}).
		Select("*").
		Find(&model)

	return model, recordtotal, recordtotalfiltered, nil
}

func MapCategory(value string) string {
	var mp = map[string]string{
		"id":          "categories.id",
		"name":        "name",
		"description": "description",
		"created_by":  "created_by",
		"updated_by":  "updated_by",
		"created_at":  "created_at",
		"updated_at":  "updated_at",
	}
	return mp[value]
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.Table("categories").Where("categories.id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) CreateCategory(req request.CreateCategoryRequest, uid string) error {
	query := `INSERT INTO categories (name, description, created_by, updated_by, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), null)`
	result := r.db.Exec(query, req.Name, req.Description, uid, nil)
	return result.Error
}
