package domains

import "gorm.io/gorm"

func (BajaAlumno) TableName() string {
	return "bajas_alumnos"
}

type BajaAlumno struct {
	gorm.Model
	AlumnoID       uint    `gorm:"not null"`
	PeriodoID      uint    `gorm:"not null"`
	PeriodoL       uint    `gorm:"not null"`
	CarreraID      uint    `gorm:"not null"`
	GrupoID        uint    `gorm:"not null"`
	Fecha          string  `gorm:"not null"`
	Tipo           string  `gorm:"not null"`
	TiempoAusencia *int    `gorm:"default:null;comment:'tiempo en cuatrimestres'"`
	CausaID        uint    `gorm:"not null"`
	Otra           *string `gorm:"type:text;default:null"`
	Observacion    *string `gorm:"type:text;default:null"`
}
