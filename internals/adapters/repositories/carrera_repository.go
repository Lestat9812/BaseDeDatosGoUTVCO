package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type CarreraRepository struct {
	db *gorm.DB
}

func NewCarreraRepository(storage *gorm.DB) *CarreraRepository {
	return &CarreraRepository{
		db: storage,
	}
}

func (r *CarreraRepository) SaveCarrera(carrera *domains.Carrera) error {
	result := r.db.Create(&carrera)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CarreraRepository) AllCarreras() ([]*domains.Carrera, error) {
	var results []*domains.Carrera

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

func (r *CarreraRepository) FindCarreraById(id int) (*domains.Carrera, error) {
	var carrera domains.Carrera
	result := r.db.
		First(&carrera, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Carrera{}, errors.New("data not found")
	}
	return &carrera, nil
}
func (r *CarreraRepository) UpdateCarrera(carrera *domains.Carrera) error {
	result := r.db.Model(&domains.Carrera{}).Where("id = ?", carrera.ID).Updates(carrera)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("carrera no actualizado")
	}
	return nil
}

func (r *CarreraRepository) DeleteCarrera(id int) error {
	var deletedCarrera domains.Carrera
	result := r.db.Where("id = ?", id).Delete(&deletedCarrera)
	if result.RowsAffected == 0 {
		return errors.New("carrera no borrado")
	}
	return nil
}
