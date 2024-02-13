package domains

import (
	"gorm.io/gorm"
)

type Grupo struct {
	// gorm.Model
	// Nombre       string
	// Cuatrimestre int
	// CarreraID    int
	// Carrera      Carrera
	// PeriodoID    int
	// Periodo      Periodo
	// Cupo         int
	// PersonalID   int
	// Personal     Personal

	gorm.Model
	Nombre       string   `gorm:"not null"`
	Cuatrimestre int      `gorm:"not null"`
	CarreraID    uint     `gorm:"not null"`
	Carrera      Carrera  `gorm:"foreignKey:CarreraID"`
	PeriodoID    uint     `gorm:"not null"`
	Periodo      Periodo  `gorm:"foreignKey:PeriodoID"`
	Cupo         int      `gorm:"not null"`
	PersonalID   uint     `gorm:"not null"`
	Personal     Personal `gorm:"foreignKey:PersonalID"`
	// Alumnos      []Alumno `gorm:"many2many:grupos;"`
}
