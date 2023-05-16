package material

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(material Material) (Material, error)
	Update(ID int, quantity int) (Material, error)
	Get() ([]Material, error)
	FindByID(ID int) (Material, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(material Material) (Material, error) {
	err := r.db.Create(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) Update(ID int, quantity int) (Material, error) {
	var material Material
	err := r.db.Model(&material).Where("id = ?", ID).Update("quantity", quantity).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) Get() ([]Material, error) {
	var materials []Material
	err := r.db.Order("updated_at DESC").Find(&materials).Error
	if err != nil {
		return materials, err
	}

	return materials, nil
}

func (r *repository) FindByID(ID int) (Material, error) {
	var material Material
	err := r.db.Where("id = ?", ID).Find(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}
