package domains

import (
	"gorm.io/gorm"
)

type Grupo struct {
	gorm.Model
	Nombre       string
	Cuatrimestre int
	CarreraID    int
	Carrera      Carrera
	PeriodoID    int
	Periodo      Periodo
	Cupo         int
	PersonalID   int
	Personal     Personal
	// Alumnos      []Alumno `gorm:"many2many:alumnos_grupos;"`
}
