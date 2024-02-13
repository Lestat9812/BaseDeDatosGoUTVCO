package domains

import (
	"time"

	"gorm.io/gorm"
)

type Alumno struct { //Usuarios de plataforma?
	gorm.Model
	Matricula        string `gorm:"column:matricula"`
	Estado           string `gorm:"column:estado;comment:(baja_temporal, baja_definitiva, activo)"`
	CarreraID        uint   `gorm:"column:id_carrera"`
	Carrera          Carrera
	Password         string `gorm:"column:password"`
	PasswordHash     string
	FechaInscripcion *time.Time `gorm:"column:fecha_inscripcion"`
	Terminos         int        `gorm:"column:terminos"`
	VisitaBiblioteca int        `gorm:"column:visitabiblioteca"`
	VisitaReglamento int        `gorm:"column:visitareglamento"`
	VisitaEstadia    int        `gorm:"column:visitaestadia"`
	FechaBiblioteca  time.Time  `gorm:"column:fechabiblioteca"`
	FechaReglamento  time.Time  `gorm:"column:fechareglamento"`
	FechaEstadia     time.Time  `gorm:"column:fechaestadia"`
	VisitaAcoso      int        `gorm:"column:visitaacoso"`
	FechaAcoso       time.Time  `gorm:"column:fechaacoso"`
	VisitaIgualdad   int        `gorm:"column:visitaigualdad"`
	FechaIgualdad    time.Time  `gorm:"column:fechaigualdad"`
	FechaClases      time.Time  `gorm:"column:fechaclases"`
	VisitaClases     int        `gorm:"column:visitaclases"`
	Grupos           []Grupo    `gorm:"many2many:alumno_grupos;"`
}

type AlumnoGrupo struct {
	AlumnoID    int    `gorm:"primaryKey"`
	GrupoID     int    `gorm:"primaryKey"`
	Estado      string `gorm:"column:estado;comment:Indica el estado del grupo para el alumno(aprobado/reprobado)"`
	Observacion string
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (AlumnoGrupo) TableName() string {
	return "alumno_grupos"
}
