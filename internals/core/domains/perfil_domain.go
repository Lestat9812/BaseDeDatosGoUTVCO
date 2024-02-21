package domains

import (
	"gorm.io/gorm"
)

func (Perfil) TableName() string {
	return "perfiles"
}

type Perfil struct {
	gorm.Model
	Perfil string
}
