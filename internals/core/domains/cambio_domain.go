package domains

import "gorm.io/gorm"

type Cambio struct {
	gorm.Model
	CalificacionID uint
	Calificacion   int
}
