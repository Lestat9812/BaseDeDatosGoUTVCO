package domains

import (
	"gorm.io/gorm"
)

type Carrera struct {
	gorm.Model
	Nombre      string
	Tipo        string
	UtcampusID  int
	Utcampus    Utcampus
	Duracion    int
	Descripcion string
	Estado      bool
	Admision    bool
	Siglas      string
	Nummax1     int
	Nummax2     int
	Nummax3     int
	DirectorID  int
	// Director Director   //no existe
}
