package domains

import (
	"time"

	"gorm.io/gorm"
)

func (ExamenProfesional) TableName() string {
	return "examen_profesional"
}

type ExamenProfesional struct {
	gorm.Model
	AlumnoID      uint      `gorm:"not null"`
	Matricula     string    `gorm:"not null"`
	Carrera       string    `gorm:"not null"`
	Generacion    uint      `gorm:"not null"`
	Correo        string    `gorm:"not null"`
	Tel           string    `gorm:"not null"`
	Memoria       string    `gorm:"not null"`
	Empresa       string    `gorm:"not null"`
	Presidente    string    `gorm:"not null"`
	Secretario    string    `gorm:"not null"`
	Vocal         string    `gorm:"not null"`
	Fecha1        time.Time `gorm:"not null"`
	Fecha2        time.Time `gorm:"not null"`
	Hora1         string    `gorm:"not null"`
	Hora2         string    `gorm:"not null"`
	Fe1           uint      `gorm:"not null"`
	Fe2           uint      `gorm:"not null"`
	FechaSugerida time.Time `gorm:"not null"`
	HSuge         string    `gorm:"not null"`
	PeriodoID     uint      `gorm:"not null"`
	Estadia       uint      `gorm:"not null"`
	Ingles        uint      `gorm:"not null"`
	Portada       uint      `gorm:"not null"`
	Oficio        uint      `gorm:"not null"`
	Pago          uint      `gorm:"not null"`
	DocCedula     uint      `gorm:"not null"`
	Status        uint      `gorm:"not null"`
	Sala          string    `gorm:"not null"`
	Observaciones string    `gorm:"not null"`
	EstadoExa     uint      `gorm:"not null"`
}
