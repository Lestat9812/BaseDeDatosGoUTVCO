package domains

import (
	"time"

	"gorm.io/gorm"
)

type Alumno struct { //Usuarios de plataforma?
	gorm.Model
	Matricula        string     `gorm:"not null;unique"`
	Estado           string     `gorm:"not null;comment:'(baja_temporal, baja_definitiva, activo)'"`
	Password         *string    `gorm:"default:null"`
	PasswordHash     string     `gorm:"-"`
	FechaInscripcion *time.Time `gorm:"default:null"`
	Terminos         int        `gorm:"not null"`
	VisitaBiblioteca int        `gorm:"not null"`
	VisitaReglamento int        `gorm:"not null"`
	VisitaEstadia    int        `gorm:"not null"`
	FechaBiblioteca  time.Time  `gorm:"not null"`
	FechaReglamento  time.Time  `gorm:"not null"`
	FechaEstadia     time.Time  `gorm:"not null"`
	VisitaAcoso      int        `gorm:"not null"`
	FechaAcoso       time.Time  `gorm:"not null"`
	VisitaIgualdad   int        `gorm:"not null"`
	FechaIgualdad    time.Time  `gorm:"not null"`
	FechaClases      time.Time  `gorm:"not null"`
	VisitaClases     int        `gorm:"not null"`
	Grupos           []Grupo    `gorm:"many2many:alumno_grupos;"`
	CarreraID        uint
	Carrera          Carrera
	EstudianteID     uint
}

type AlumnoGrupo struct {
	AlumnoID    int    `gorm:"foreignKey"`
	GrupoID     int    `gorm:"foreignKey"`
	Estado      string `gorm:"column:estado;comment:Indica el estado del grupo para el alumno(aprobado/reprobado)"`
	Observacion string
	gorm.Model
}

func (AlumnoGrupo) TableName() string {
	return "alumno_grupos"
}
