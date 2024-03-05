package domains

import "gorm.io/gorm"

type Estadia struct {
	gorm.Model
	GrupoMateriaID    uint     `gorm:"not null"`
	AlumnoID          uint     `gorm:"not null"`
	Proyecto          string   `gorm:"default:null;comment:'Nombre del Proyecto de Estadia'"`
	PrimeroInter      *float64 `gorm:"default:null"`
	SegundoInter      *float64 `gorm:"default:null"`
	TerceroInter      *float64 `gorm:"default:null"`
	AsesorEmpresarial *string  `gorm:"default:null"`
	Empresa           *string  `gorm:"default:null"`
	Liberado          *string  `gorm:"default:null"`
	Nivel             string   `gorm:"not null"`
	CarreraID         uint     `gorm:"not null"`
}
