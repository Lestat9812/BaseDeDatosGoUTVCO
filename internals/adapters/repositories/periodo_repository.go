package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type PeriodoRepository struct {
	db *gorm.DB
}

func NewPeriodoRepository(storage *gorm.DB) *PeriodoRepository {
	return &PeriodoRepository{
		db: storage,
	}
}

func (r *PeriodoRepository) SavePeriodo(periodo *domains.Periodo) error {
	result := r.db.Create(&periodo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PeriodoRepository) AllPeriodos() ([]*domains.Periodo, error) {
	var results []*domains.Periodo

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

func (r *PeriodoRepository) FindPeriodoById(id int) (*domains.Periodo, error) {
	var periodo domains.Periodo
	result := r.db.
		First(&periodo, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Periodo{}, errors.New("data not found")
	}
	return &periodo, nil
}
func (r *PeriodoRepository) UpdatePeriodo(periodo *domains.Periodo) error {
	result := r.db.Model(&domains.Periodo{}).Where("id = ?", periodo.ID).Updates(periodo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("periodo no actualizado")
	}
	return nil
}

func (r *PeriodoRepository) DeletePeriodo(id int) error {
	var deletedPeriodo domains.Periodo
	result := r.db.Where("id = ?", id).Delete(&deletedPeriodo)
	if result.RowsAffected == 0 {
		return errors.New("periodo no borrado")
	}
	return nil
}
