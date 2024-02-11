package domains

import (
	"time"

	"gorm.io/gorm"
)

type Periodo struct {
	gorm.Model
	InicioAdmision *time.Time `gorm:"column:inicio_admision"`
	FinAdmision    *time.Time `gorm:"column:fin_admision"`
	FechaInicio    time.Time  `gorm:"column:fecha_inicio"`
	FechaFin       time.Time  `gorm:"column:fecha_fin"`
	InicioReins    *time.Time `gorm:"column:inicio_reins"`
	FinReins       *time.Time `gorm:"column:fin_reins"`
	ExamenProf     int        `gorm:"column:examenprof"`
	Titulacion     int        `gorm:"column:titulacion"`
	TitulacionIng  int        `gorm:"column:titulacioning"`
	Primero        bool       `gorm:"column:primero;comment:activar/desactivar primer parcial para calificaciones"`
	Segundo        bool       `gorm:"column:segundo;comment:activar/desactivar segundo parcial para calificaciones"`
	Tercero        bool       `gorm:"column:tercero;comment:activar/desactivar tercer parcial para calificaciones"`
	P1             int        `gorm:"column:p1"`
	P2             int        `gorm:"column:p2"`
	P3             int        `gorm:"column:p3"`
	InicioEv       *time.Time `gorm:"column:inicio_ev"`
	FinEv          *time.Time `gorm:"column:fin_ev"`
	Parcial        int        `gorm:"column:parcial"`
}
