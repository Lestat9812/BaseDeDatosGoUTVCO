package domains

import (
	"gorm.io/gorm"
)

type Carrera struct {
	gorm.Model
	Nombre      string  `gorm:"not null"`
	Tipo        string  `gorm:"not null"`
	Duracion    int     `gorm:"not null"`
	Descripcion *string `gorm:"default:null"`
	Estado      bool    `gorm:"not null"`
	Admision    bool    `gorm:"not null"`
	Siglas      *string `gorm:"default:null"`
	Nummax1     int     `gorm:"not null"`
	Nummax2     int     `gorm:"not null"`
	Nummax3     int     `gorm:"not null"`
	UtcampusID  uint
	Utcampus    Utcampus
	PersonalID  uint
	Personal    Personal //Director
	//DirectorID  int
	// Director Director   //no existe
}
