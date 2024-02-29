package domains

import "gorm.io/gorm"

type Aspirante struct {
	gorm.Model
	Fecha        string `gorm:"not null;comment:'fecha de solicitud de admision'"`
	Folio        int    `gorm:"not null;comment:'folio consecutivo por periodo'"`
	Convocatoria int    `gorm:"not null"`
	Carrera1ID   uint   `gorm:"not null;comment:'primera opcion de carrera'"` //No es un índice
	Carrera2ID   *uint  `gorm:"default:null;comment:'segunda opcion de carrera'"`
	CarFinal     *int   `gorm:"default:null"`
	PeriodoID    uint   `gorm:"not null;comment:'periodo en el que solicito admision'"`
	Nivel        int    `gorm:"not null;comment:'nivel del proceso de admision'"`
	Password     string `gorm:"not null"` //Necesario?
	PasswordHash string `gorm:"-"`
	Opcion       int    `gorm:"not null"`
	OpPor        string `gorm:"not null"`
	PInscripcion int    `gorm:"not null"`
	PCred        *int   `gorm:"default:null"`
	PPlayera     *int   `gorm:"default:null"`
	PUnigastro   *int   `gorm:"default:null"`
	PFilipina    *int   `gorm:"default:null"`
	PrefichaID   uint
	Estudiante   Estudiante
}
