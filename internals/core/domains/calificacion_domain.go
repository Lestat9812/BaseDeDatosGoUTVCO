package domains

import "gorm.io/gorm"

func (Calificacion) TableName() string {
	return "calificaciones"
}

type Calificacion struct {
	gorm.Model
	GrupoMateriaID uint     `gorm:"not null"`
	AlumnoID       uint     `gorm:"not null"`
	Primero        *float64 `gorm:"default:null"`
	Segundo        *float64 `gorm:"default:null"`
	Tercero        *float64 `gorm:"default:null"`
	Cambios        []Cambio
	Extraordinario Extraordinario
	Recursamiento  Recursamiento
}
