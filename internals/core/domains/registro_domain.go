package domains

import "gorm.io/gorm"

type Registro struct {
	gorm.Model
	TipoID           uint   `gorm:"not null"` //tipo de qué????
	GrupoMateriaID   uint   `gorm:"default:null"`
	NumDoc           uint   `gorm:"default:null"`
	GrupoID          uint   `gorm:"default:null"`
	EstadoRegistroID uint   `gorm:"not null"`
	Extension        string `gorm:"size:5;not null"`
	Observaciones    string `gorm:"type:text;not null"`
	PersonalID       uint   `gorm:"not null"` //Persona que hizo el registro??
}

type RegVistos struct { //many to many???
	DocumentoID int
	PersonalID  int
}
