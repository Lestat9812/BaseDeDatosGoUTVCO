package domains

import "gorm.io/gorm"

func (BecaAlumno) TableName() string {
	return "becas_alumnos"
}

type BecaAlumno struct {
	gorm.Model
	AlumnoID    uint   `gorm:"primaryKey"`
	TipoID      uint   `gorm:"primaryKey"`
	FechaInicio string `gorm:"not null"`
	FechaFin    string `gorm:"not null"`
	Estado      bool   `gorm:"not null"`
}
