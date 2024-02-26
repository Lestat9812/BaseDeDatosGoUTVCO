package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type EstudianteRepository struct {
	db *gorm.DB
}

func NewEstudianteRepository(storage *gorm.DB) *EstudianteRepository {
	return &EstudianteRepository{
		db: storage,
	}
}

func (r *EstudianteRepository) SaveEstudiante(estudiante *domains.Estudiante) error {
	result := r.db.Create(&estudiante)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *EstudianteRepository) AllEstudiantes() ([]*domains.Estudiante, error) {
	var results []*domains.Estudiante

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

func (r *EstudianteRepository) FindEstudianteById(id int) (*domains.Estudiante, error) {
	var estudiante domains.Estudiante
	result := r.db.
		First(&estudiante, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Estudiante{}, errors.New("data not found")
	}
	return &estudiante, nil
}
func (r *EstudianteRepository) UpdateEstudiante(estudiante *domains.Estudiante) error {
	result := r.db.Model(&domains.Estudiante{}).Where("id = ?", estudiante.ID).Updates(estudiante)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("estudiante no actualizado")
	}
	return nil
}

func (r *EstudianteRepository) DeleteEstudiante(id int) error {
	var deletedEstudiante domains.Estudiante
	result := r.db.Where("id = ?", id).Delete(&deletedEstudiante)
	if result.RowsAffected == 0 {
		return errors.New("estudiante no borrado")
	}
	return nil
}
