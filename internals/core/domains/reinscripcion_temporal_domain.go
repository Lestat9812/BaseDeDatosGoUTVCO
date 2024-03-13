package domains

import (
	"time"

	"gorm.io/gorm"
)

func (ReinscripcionTemporal) TableName() string {
	return "reinscripciontemporal"
}

type ReinscripcionTemporal struct {
	gorm.Model
	AlumnoID           uint      `gorm:"not null"`
	CarreraID          uint      `gorm:"not null"`
	PeriodoID          uint      `gorm:"not null"`
	Carrera            string    `gorm:"not null;size:150"`
	Matricula          string    `gorm:"not null;size:15"`
	Nombre             string    `gorm:"not null;size:25"`
	ApellidoP          string    `gorm:"not null;size:20"`
	ApellidoM          string    `gorm:"not null;size:20"`
	Edad               uint      `gorm:"not null"`
	Sexo               string    `gorm:"not null;size:5"`
	Cuatrimestre       uint      `gorm:"not null"`
	Calle              string    `gorm:"not null;size:50"`
	NumInt             uint      `gorm:"not null"`
	Localidad          string    `gorm:"not null;size:200"`
	Region             string    `gorm:"not null;size:200"`
	CURP               string    `gorm:"not null;size:50"`
	Municipio          string    `gorm:"not null;size:50"`
	Distrito           string    `gorm:"not null;size:200"`
	Estado             string    `gorm:"not null;size:20"`
	Telefono           string    `gorm:"not null;size:15"`
	TelefEmergencia    string    `gorm:"not null;size:15"`
	Discapacidad       string    `gorm:"not null;size:50"`
	OtraDiscapacidad   string    `gorm:"size:30"`
	LenguaM            string    `gorm:"size:30"`
	NombreTutor        string    `gorm:"not null;size:50"`
	DireccionTutor     string    `gorm:"not null;size:80"`
	TelefonoTutor      string    `gorm:"not null;size:15"`
	Fecha              time.Time `gorm:"not null"`
	Observaciones      string    `gorm:"size:200"`
	EstadoRegistroID   uint      `gorm:"not null"` //??
	FirmaCarta         string    `gorm:"not null;size:255"`
	CartaCompromiso    uint      `gorm:"not null"`
	INE                uint      `gorm:"not null"`
	ObservacionesINE   string    `gorm:"not null;size:255"`
	ObservacionesCarta string    `gorm:"not null;size:255"`
}
