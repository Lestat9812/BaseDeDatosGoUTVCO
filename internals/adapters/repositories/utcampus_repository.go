package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type UtcampusRepository struct {
	db *gorm.DB
}

func NewUtcampusRepository(storage *gorm.DB) *UtcampusRepository {
	return &UtcampusRepository{
		db: storage,
	}
}

func (r *UtcampusRepository) SaveUtcampus(utcampus *domains.Utcampus) error {
	result := r.db.Create(&utcampus)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UtcampusRepository) AllUtcampus() ([]*domains.Utcampus, error) {
	var results []*domains.Utcampus

	if res := r.db.
		Find(&results); res.Error != nil {
		return nil, res.Error
	}

	if len(results) == 0 {
		return nil, errors.New("data not found")
	} else {
		return results, nil
	}
}

func (r *UtcampusRepository) FindUtcampusById(id int) (*domains.Utcampus, error) {
	var utcampus domains.Utcampus
	result := r.db.
		First(&utcampus, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Utcampus{}, errors.New("data not found")
	}
	return &utcampus, nil
}
func (r *UtcampusRepository) UpdateUtcampus(utcampus *domains.Utcampus) error {
	result := r.db.Model(&domains.Utcampus{}).Where("id = ?", utcampus.ID).Updates(utcampus)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("utcampus no actualizado")
	}
	return nil
}

func (r *UtcampusRepository) DeleteUtcampus(id int) error {
	var deletedUtcampus domains.Utcampus
	result := r.db.Where("id = ?", id).Delete(&deletedUtcampus)
	if result.RowsAffected == 0 {
		return errors.New("utcampus no borrado")
	}
	return nil
}
