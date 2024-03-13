package domains

import "gorm.io/gorm"

func (PagosFicha) TableName() string {
	return "pagosficha"
}

type PagosFicha struct {
	gorm.Model
	PrefichaID    uint   `gorm:"not null"`
	EstudianteID  uint   `gorm:"not null"`
	PeriodoID     uint   `gorm:"not null"`
	Fecha         string `gorm:"not null"`
	Extension     string `gorm:"not null"`
	NumReferencia uint   `gorm:"not null"`
	Status        uint   `gorm:"not null"`
	Observaciones string `gorm:"not null"`
}
