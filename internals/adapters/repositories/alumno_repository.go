package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"golang.org/x/crypto/bcrypt"
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(alumno.Password), 10)
	if err != nil {
		return errors.New(err.Error())
	} else {
		alumno.Password = string(passwordHash)
	}
	result := r.db.Create(&alumno)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AlumnoRepository) AllAlumnos() ([]*domains.AlumnoSinNada, error) {
	var al []*domains.AlumnoSinNada
	result2 := r.db.Model(domains.Alumno{}).Select("alumnos.id, alumnos.matricula, preficha.nombre, preficha.apellido_p, preficha.apellido_m, alumnos.carrera_id").Joins("left join estudiantes on estudiantes.id = alumnos.estudiante_id").Joins("left join preficha on preficha.id=estudiantes.preficha_id").Scan(&al)

	if result2.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return al, nil
}

func (r *AlumnoRepository) FindAlumnoById(id int) (*domains.Alumno, error) {
	var alumno domains.Alumno
	result := r.db.Preload(clause.Associations).First(&alumno)
	// result := r.db.Preload(clause.Associations).Preload("Grupos.Carrera.Personal.Perfiles").First(&alumno)

	if result.RowsAffected == 0 {
		return &domains.Alumno{}, errors.New("data not found")
	}
	return &alumno, nil
}

func (r *AlumnoRepository) FindAlumnoGrupoById(id int) (*domains.Alumno, error) {
	var alumno domains.Alumno
	result := r.db.Preload("Grupos").First(&alumno)

	if result.RowsAffected == 0 {
		return &domains.Alumno{}, errors.New("data not found")
	}
	return &alumno, nil
}

func (r *AlumnoRepository) FindAlumnoCalificacionesById(id int) (*domains.Alumno, error) {
	var alumno domains.Alumno
	result := r.db.Preload("Calificaciones").First(&alumno)

	if result.RowsAffected == 0 {
		return &domains.Alumno{}, errors.New("data not found")
	}
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
