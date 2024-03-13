package domains

import (
	"gorm.io/gorm"
)

// TableName overrides the table name used by User to `profiles`
func (Personal) TableName() string {
	return "personal"
}

type Personal struct {
	gorm.Model
	Grado          string   `gorm:"column:grado"`
	Nombre         string   `gorm:"column:nombre"`
	ApellidoP      string   `gorm:"column:apellido_p"`
	ApellidoM      string   `gorm:"column:apellido_m"`
	User           string   `gorm:"column:user"`
	Password       string   `gorm:"column:password"`
	PasswordHash   string   `gorm:"-"`
	Estado         bool     `gorm:"column:estado"`
	Adscrito       uint     `gorm:"column:adscrito"`
	Perfiles       []Perfil `gorm:"many2many:usuarios_perfiles;"`
	Autoevaluacion Autoevaluacion
	// DetalleAutoevaluacion DetalleAutoevaluacion
	Curriculum         Curriculum
	Documento          Documento
	Evaluacion         Evaluacion
	EvaluacionDirector EvaluacionDirector
	Recursamiento      Recursamiento
	Registro           Registro
	// RegVistos          RegVistos

	// GrupoMateria []GrupoMateria //Así no era
	// PerfilID  uint   `gorm:"column:id_perfil;index"` //Se pueden más de uno
}
