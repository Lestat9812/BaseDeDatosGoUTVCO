package domains

import (
	"time"

	"gorm.io/gorm"
)

// Estudiante represents the Estudiante model
type Estudiante struct {
	gorm.Model
	// Personal Information
	Nombre                 string    `gorm:"not null;comment:'Nombre del aspirante'"`
	ApellidoP              string    `gorm:"not null;comment:'Apellido paterno del aspirante'"`
	ApellidoM              *string   `gorm:"default:null;comment:'Apellido materno del aspirante'"`
	Sexo                   string    `gorm:"not null;comment:'Género del aspirante (M/F)'"`
	Discapacidad           *string   `gorm:"default:null;comment:'Discapacidad del aspirante'"`
	Sangre                 *string   `gorm:"default:null;comment:'Tipo de sangre del aspirante'"`
	Idioma                 string    `gorm:"not null;comment:'Lengua materna del aspirante'"`
	FechaNacimiento        time.Time `gorm:"not null;comment:'Fecha de nacimiento del aspirante'"`
	Edad                   *int      `gorm:"default:null;comment:'Edad del aspirante'"`
	EstadoCivil            string    `gorm:"not null;comment:'Estado civil del aspirante (S/C)'"`
	Curp                   string    `gorm:"not null;comment:'CURP del aspirante'"`
	Nss                    string    `gorm:"not null;comment:'Número de Seguro Social (NSS)'"`
	Calle                  string    `gorm:"not null;comment:'Calle de la dirección del aspirante'"`
	Numero                 *string   `gorm:"default:null;comment:'Número del domicilio del aspirante'"`
	Colonia                *string   `gorm:"default:null;comment:'Colonia donde vive el aspirante'"`
	Distrito               string    `gorm:"not null;comment:'Distrito del aspirante'"`
	Ciudad                 *string   `gorm:"default:null;comment:'Ciudad reconocida del aspirante'"`
	Estado                 string    `gorm:"not null;comment:'Entidad federativa del aspirante'"`
	Cp                     int       `gorm:"not null;comment:'Código Postal (CP) del aspirante'"`
	Localidad              *string   `gorm:"default:null;comment:'Localidad del aspirante'"`
	Municipio              string    `gorm:"not null;comment:'Municipio donde reside el aspirante'"`
	Email                  *string   `gorm:"default:null;comment:'Correo electrónico del aspirante'"`
	Telefono               *string   `gorm:"default:null;comment:'Número telefónico del aspirante'"`
	EscProcedencia         string    `gorm:"not null;comment:'Nombre de la escuela de procedencia'"`
	SubsistemaProcedencia  string    `gorm:"not null;comment:'Subsistema de la escuela de procedencia'"`
	TipoProcedencia        string    `gorm:"not null;comment:'Tipo de procedencia (Privada/Pública/Cooperación)'"`
	AreaProcedencia        string    `gorm:"not null;comment:'Área de bachillerato de procedencia'"`
	FechaEgresoProcedencia time.Time `gorm:"not null;comment:'Fecha de egreso del nivel medio superior'"`
	MunicipioProcedencia   string    `gorm:"not null;comment:'Municipio donde se encuentra la escuela de procedencia'"`
	EstadoProcedencia      string    `gorm:"not null;comment:'Estado donde se encuentra la escuela de procedencia'"`
	PromedioProcedencia    *float64  `gorm:"default:null;comment:'Promedio obtenido en el nivel medio superior'"`
	DuracionProcedencia    int       `gorm:"not null;comment:'Duración del plan de estudios del nivel medio superior (1/2)'"`
	Secundaria             string    `gorm:"not null;comment:'Nombre de la secundaria donde cursó'"`
	PromedioSecundaria     *float64  `gorm:"default:null;comment:'Promedio obtenido en la secundaria'"`
	Tutor                  *string   `gorm:"default:null;comment:'Nombre del tutor'"`
	TutorAp                string    `gorm:"not null;comment:'Apellido paterno del tutor'"`
	TutorAm                *string   `gorm:"default:null;comment:'Apellido materno del tutor'"`
	DireccionTutor         *string   `gorm:"default:null;comment:'Dirección del tutor'"`
	TelefonoTutor          *string   `gorm:"default:null;comment:'Número telefónico del tutor'"`
	EmailTutor             *string   `gorm:"default:null;comment:'Correo electrónico del tutor'"`
	Observaciones          *string   `gorm:"default:null;comment:'Observaciones sobre el aspirante'"`
	Obserficha             string    `gorm:"not null;comment:'Observaciones en ficha'"`
	// IdPficha               int       `gorm:"not null;comment:'ID de la preficha'"`
	FechaFicha       time.Time `gorm:"not null;comment:'Fecha de la ficha'"`
	MatriculaCeneval string    `gorm:"not null;comment:'Matrícula del examen CENEVAL'"`
	PassFicha        string    `gorm:"not null;comment:'Contraseña de la ficha'"`
	CarFicha         int       `gorm:"not null;comment:'Carrera de la ficha'"`
	Car2Ficha        int       `gorm:"not null;comment:'Segunda opción de carrera de la ficha'"`
	NivelFi          int       `gorm:"not null;comment:'Nivel FI'"`
	OpPor            string    `gorm:"not null;comment:'OP POR'"`
	Acta             int       `gorm:"not null;comment:'Acta'"`
	Curpp            int       `gorm:"not null;comment:'CURPP'"`
	Identi           int       `gorm:"not null;comment:'Identificador'"`
	Cer              int       `gorm:"not null;comment:'CER'"`
	Curpins          int       `gorm:"not null;comment:'CURPINS'"`
	Certins          int       `gorm:"not null;comment:'CERTINS'"`
	Actains          int       `gorm:"not null;comment:'ACTAINS'"`
	Certmedicoins    int       `gorm:"not null;comment:'CERTMEDICOINS'"`
	Compdomiins      int       `gorm:"not null;comment:'COMPDOMIINS'"`
	Pagoins          int       `gorm:"not null;comment:'PAGOINS'"`
	DesCermediins    *string   `gorm:"default:null;comment:'Des Cermediins'"`
	DesCertins       *string   `gorm:"default:null;comment:'Des Certins'"`
	DesActains       *string   `gorm:"default:null;comment:'Des Actains'"`
	DesComdomiins    *string   `gorm:"default:null;comment:'Des Comdomiins'"`
	Cartacomp        int       `gorm:"not null;comment:'Carta de compromiso'"`
	DPagoins         *string   `gorm:"default:null;comment:'D Pagoins'"`
	DCertins         *string   `gorm:"default:null;comment:'D Certins'"`
	DCertmedicoins   *string   `gorm:"default:null;comment:'D Certmedicoins'"`
	DActains         *string   `gorm:"default:null;comment:'D Actains'"`
	DCompins         *string   `gorm:"default:null;comment:'D Compins'"`
	DCurp            *string   `gorm:"default:null;comment:'D Curp'"`
	PrefichaID       uint      `gorm:"not null;comment:'ID de la preficha'"`
	AspiranteID      uint
	Alumno           Alumno //`gorm:"foreignKey:Matricula"`
}
