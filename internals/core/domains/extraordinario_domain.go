package domains

import (
	"gorm.io/gorm"
)

type Extraordinario struct {
	gorm.Model
	CalificacionID uint `gorm:"not null"`
	Calificacion   *int `gorm:"default:null"`
}
