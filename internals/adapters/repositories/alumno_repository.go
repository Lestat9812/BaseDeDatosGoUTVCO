package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AlumnoRepository struct {
	db *gorm.DB
}

func NewAlumnoRepository(storage *gorm.DB) *AlumnoRepository {
	return &AlumnoRepository{
		db: storage,
	}
}

func (r *AlumnoRepository) SaveAlumno(alumno *domains.Alumno) error {
	result := r.db.Create(&alumno)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AlumnoRepository) AllAlumnos() ([]*domains.AlumnoSinNada, error) {
	// var alumno domains.Alumno
	var al []*domains.AlumnoSinNada
	// result := r.db.Preload(clause.Associations).Preload("Grupos.Carrera.Personal.Perfiles").First(&alumno)
	result2 := r.db.Model(domains.Alumno{}).Select("alumnos.id, alumnos.matricula, estudiantes.nombre").Joins("left join estudiantes on estudiantes.id = alumnos.estudiante_id").Scan(&al)

	// fmt.Print(al)
	// if result.RowsAffected == 0 {
	// 	return &AlumnoSinNada{}, errors.New("data not found")
	// 	// return &domains.Alumno{}, errors.New("data not found")
	// }
	if result2.RowsAffected == 0 {
		return nil, errors.New("data not found")
		// return &domains.Alumno{}, errors.New("data not found")
	}
	return al, nil
	// var results []*domains.Alumno

	// if res := r.db.Preload(clause.Associations).
	// 	Find(&results); res.Error != nil {
	// 	return nil, res.Error
	// }

	// if len(results) == 0 {
	// 	return nil, errors.New("data not found")
	// } else {
	// 	return results, nil
	// }
}

func (r *AlumnoRepository) FindAlumnoById(id int) (*domains.Alumno, error) {
	var alumno domains.Alumno
	// var al AlumnoSinNada
	result := r.db.Preload(clause.Associations).Preload("Grupos.Carrera.Personal.Perfiles").First(&alumno)
	// result2 := r.db.Model(domains.Alumno{}).Select("id, matricula, password").Where("id=?", id).Scan(&al)

	// fmt.Print(al)
	if result.RowsAffected == 0 {
		// 	return &AlumnoSinNada{}, errors.New("data not found")
		return &domains.Alumno{}, errors.New("data not found")
	}
	// if result2.RowsAffected == 0 {
	// return &AlumnoSinNada{}, errors.New("data not found")
	// return &domains.Alumno{}, errors.New("data not found")
	// }
	return &alumno, nil
}

func (r *AlumnoRepository) UpdateAlumno(alumno *domains.Alumno) error {
	result := r.db.Model(&domains.Alumno{}).Where("id = ?", alumno.ID).Updates(alumno)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("alumno no actualizado")
	}
	return nil
}

func (r *AlumnoRepository) DeleteAlumno(id int) error {
	var deletedAlumno domains.Alumno
	result := r.db.Where("id = ?", id).Delete(&deletedAlumno)
	if result.RowsAffected == 0 {
		return errors.New("alumno no borrado")
	}
	return nil
}
