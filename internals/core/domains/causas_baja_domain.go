package domains

import "gorm.io/gorm"

func (CausaBaja) TableName() string {
	return "causas_baja"
}

type CausaBaja struct {
	gorm.Model
	Causa      string `gorm:"not null;comment:'nombre de la causa'"`
	Estado     bool   `gorm:"not null;comment:'si es definida por el formato o es una nueva causa (otra)'"`
	BajaAlumno BajaAlumno
}
