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
	PersonalID uint
	Personal   Personal
	PerfilID   uint
	Perfil     Perfil
}
