package domains

import "gorm.io/gorm"

func (Materia) TableName() string {
	return "materias"
}

type Materia struct {
	gorm.Model
	Nombre       string `gorm:"not null"`
	Horas        int    `gorm:"default:null"`
	Cuatrimestre int    `gorm:"not null;comment:'cuarimestre de la materia'"`
	CarreraID    uint   `gorm:"not null;comment:'carrera a la que pertenece la materia'"`
	Carrera      Carrera
	Estado       *bool `gorm:"default:null"`
}
