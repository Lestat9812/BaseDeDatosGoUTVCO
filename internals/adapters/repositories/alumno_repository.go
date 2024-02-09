package repositories

import (

	//"gitlab.com/l-cm/api-giras/internals/core/domains"
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type AlumnoRepository struct {
	db *gorm.DB
}

func NewAlumnoRepository(storage *gorm.DB) *AlumnoRepository {
	return &AlumnoRepository{
		db: storage,
	}
}

func (r *AlumnoRepository) SaveAlumno(encargado *domains.Alumno) error {
	result := r.db.Create(&encargado)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AlumnoRepository) AllAlumnos() ([]*domains.Alumno, error) {
	var results []*domains.Alumno

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

func (r *AlumnoRepository) FindAlumnoById(id int) (*domains.Alumno, error) {
	var encargado domains.Alumno
	result := r.db.
		First(&encargado, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Alumno{}, errors.New("data not found")
	}
	return &encargado, nil
}
func (r *AlumnoRepository) UpdateAlumno(encargado *domains.Alumno) error {
	result := r.db.Model(&domains.Alumno{}).Where("id = ?", encargado.ID).Updates(encargado)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("encargado no actualizado")
	}
	return nil
}

func (r *AlumnoRepository) DeleteAlumno(id int) error {
	var deletedAlumno domains.Alumno
	result := r.db.Where("id = ?", id).Delete(&deletedAlumno)
	if result.RowsAffected == 0 {
		return errors.New("encargado no borrado")
	}
	return nil
}
