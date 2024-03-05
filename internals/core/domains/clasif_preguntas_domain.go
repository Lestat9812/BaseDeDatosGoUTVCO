package domains

import "gorm.io/gorm"

func (ClasifPreg) TableName() string {
	return "clasif_preg"
}

type ClasifPreg struct {
	gorm.Model
	Clasificacion string `gorm:"not null"`
}
