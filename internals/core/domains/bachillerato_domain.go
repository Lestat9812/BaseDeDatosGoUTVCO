package domains

import "gorm.io/gorm"

type Bachillerato struct {
	gorm.Model
	Clave     string `gorm:"not null"`
	NombreBA  string `gorm:"not null"`
	TipoProc  string `gorm:"not null"`
	EstadoPro string `gorm:"not null"`
}
