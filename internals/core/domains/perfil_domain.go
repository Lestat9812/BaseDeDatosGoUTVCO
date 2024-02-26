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

type UsuariosPerfiles struct {
	gorm.Model
	PersonalID uint     `gorm:"primaryKey"`
	Personal   Personal //muestra al personal
	PerfilID   uint     `gorm:"primaryKey"`
	Perfil     Perfil   //muestra los perfiles
}
