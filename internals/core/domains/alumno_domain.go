package domains

import (
	"time"

	"gorm.io/gorm"
)

type Alumno struct { //Usuarios de plataforma?
	gorm.Model
	Matricula        string     `gorm:"not null;unique"`
	Estado           string     `gorm:"not null;comment:'(baja_temporal, baja_definitiva, activo)'"`
	Password         string     //`gorm:"default:null"`
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
	EstudianteID     uint       //Relación tiene uno
	Grupos           []Grupo    `gorm:"many2many:alumnos_grupos;"` //Relación many to many
	CarreraID        uint
	Carrera          Carrera //Relación pertenece a
	Calificaciones   []Calificacion
	// Perfiles         []Perfil //Relación has many  //Necesita perfil?
}

type AlumnoGrupo struct {
	gorm.Model
	AlumnoID    int    `gorm:"primaryKey"`
	GrupoID     int    `gorm:"primaryKey"`
	Estado      string `gorm:"column:estado;comment:Indica el estado del grupo para el alumno(aprobado/reprobado)"`
	Observacion string
}

// func (AlumnoGrupo) TableName() string {
// 	return "alumno_grupos"
// }
