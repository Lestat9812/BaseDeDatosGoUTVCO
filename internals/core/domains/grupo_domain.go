package domains

import (
	"gorm.io/gorm"
)

type Grupo struct {
	gorm.Model
	Nombre        string `gorm:"not null"`
	Cuatrimestre  int    `gorm:"not null"`
	CarreraID     uint   `gorm:"not null"`
	Carrera       Carrera
	PeriodoID     uint `gorm:"not null"`
	Periodo       Periodo
	Cupo          int  `gorm:"not null"`
	PersonalID    uint `gorm:"not null"` //TUTOR?
	Personal      Personal
	GrupoMaterias []GrupoMateria
}

func (GrupoMateria) TableName() string {
	return "grupos_materias"
}

type GrupoMateria struct {
	gorm.Model
	GrupoID uint
	// Grupo
	MateriaID      uint
	Materia        Materia
	PersonalID     uint //Maestro?
	Personal       Personal
	Primero        int
	Segundo        int
	Tercero        int
	Acta1          int
	Acta2          int
	Acta3          int
	Calificaciones []Calificacion
}
