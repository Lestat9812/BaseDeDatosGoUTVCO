package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type MateriaRepository struct {
	db *gorm.DB
}

func NewMateriaRepository(storage *gorm.DB) *MateriaRepository {
	return &MateriaRepository{
		db: storage,
	}
}

func (r *MateriaRepository) SaveMateria(materia *domains.Materia) error {
	result := r.db.Create(&materia)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MateriaRepository) AllMaterias() ([]*domains.Materia, error) {
	var results []*domains.Materia

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

func (r *MateriaRepository) FindMateriaById(id int) (*domains.Materia, error) {
	var materia domains.Materia
	result := r.db.
		First(&materia, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Materia{}, errors.New("data not found")
	}
	return &materia, nil
}
func (r *MateriaRepository) UpdateMateria(materia *domains.Materia) error {
	result := r.db.Model(&domains.Materia{}).Where("id = ?", materia.ID).Updates(materia)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("materia no actualizado")
	}
	return nil
}

func (r *MateriaRepository) DeleteMateria(id int) error {
	var deletedMateria domains.Materia
	result := r.db.Where("id = ?", id).Delete(&deletedMateria)
	if result.RowsAffected == 0 {
		return errors.New("materia no borrado")
	}
	return nil
}
