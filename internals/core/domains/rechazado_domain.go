package domains

import "gorm.io/gorm"

type Rechazado struct {
	gorm.Model
	DocumentoID uint `gorm:"not null"`
}
