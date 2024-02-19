package domains

import (
	"gorm.io/gorm"
)

type Grupo struct {
	gorm.Model
	Nombre       string    `gorm:"not null"`
	Cuatrimestre int       `gorm:"not null"`
	CarreraID    uint      `gorm:"not null"`
	Carrera      Carrera   `gorm:"foreignKey:CarreraID"`
	PeriodoID    uint      `gorm:"not null"`
	Periodo      Periodo   `gorm:"foreignKey:PeriodoID"`
	Cupo         int       `gorm:"not null"`
	PersonalID   uint      `gorm:"not null"`
	Personal     Personal  `gorm:"foreignKey:PersonalID"`
	Materias     []Materia `gorm:"manytomany:grupos_materias;"`
}

type GrupoMateria struct {
	gorm.Model
	GrupoID   int `gorm:"primaryKey"`
	MateriaID int `gorm:"primaryKey"`
	// PersonalID int `gorm:"primaryKey"`
	Primero int
	Segundo int
	Tercero int
	Acta1   int
	Acta2   int
	Acta3   int
}
