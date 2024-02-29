package domains

import "gorm.io/gorm"

type Preficha struct {
	gorm.Model
	Nombre                 string   `gorm:"not null;comment:'nomnre del aspirante'"`
	ApellidoP              string   `gorm:"not null;comment:'apellido paterno de aspirante'"`
	ApellidoM              *string  `gorm:"default:null;comment:'apellido materno del aspirante'"`
	Sexo                   string   `gorm:"not null;comment:'genero del aspitante (m/f)'"`
	Discapacidad           *string  `gorm:"default:null"`
	Sangre                 *string  `gorm:"default:null;comment:'tipo de sangre'"`
	Idioma                 string   `gorm:"not null;comment:'lengua materna'"`
	FechaNacimiento        string   `gorm:"not null;comment:'fecha de nacimiento del aspirante'"`
	Edad                   *int     `gorm:"default:null"`
	EstadoCivil            string   `gorm:"not null;comment:'estado civil del aspirante (s/c)'"`
	CURP                   string   `gorm:"not null;comment:'CURP del aspirante'"`
	NSS                    string   `gorm:"not null"`
	Calle                  string   `gorm:"not null;comment:'Calle de la direccion del aspirante'"`
	Numero                 *string  `gorm:"default:null;comment:'numero del domicilio del aspirante'"`
	Colonia                *string  `gorm:"default:null;comment:'colonia donde vive el aspirante'"`
	Distrito               string   `gorm:"not null"`
	Ciudad                 *string  `gorm:"default:null;comment:'ciudad reconocida del aspirante'"`
	Estado                 string   `gorm:"not null;comment:'entidad federativa'"`
	CP                     int      `gorm:"not null;comment:'codigo postal del aspirante'"`
	Localidad              *string  `gorm:"default:null;comment:'localidad del aspirante'"`
	Municipio              string   `gorm:"not null;comment:'municipio donde reside el aspirante'"`
	Email                  *string  `gorm:"default:null;comment:'correo electronico del aspirante'"`
	Telefono               *string  `gorm:"default:null;comment:'numero telefonico del aspirante'"`
	EscuelaProcedencia     string   `gorm:"not null;comment:'nombre de la escuela de procedencia'"`
	SubsistemaProcedencia  string   `gorm:"not null;comment:'subsistema de la escuela de procedencia'"`
	TipoProcedencia        string   `gorm:"not null;comment:'(privada, publica, por cooperacion)'"`
	AreaProcedencia        string   `gorm:"not null;comment:'area de bachillerato de procedencia'"`
	FechaEgresoProcedencia string   `gorm:"not null;comment:'fecha de egreso del nivel medio superior'"`
	MunicipioProcedencia   string   `gorm:"not null;comment:'municipio donde se encuentra la escuela de procedencia'"`
	EstadoProcedencia      string   `gorm:"not null"`
	PromedioProcedencia    *float64 `gorm:"default:null;comment:'promedio obtenido en el nivel medio superior'"`
	DuracionProcedencia    int      `gorm:"not null;comment:'duracion del plan de estudios del nivel medio superior (1/2)'"`
	Secundaria             string   `gorm:"not null;comment:'nombre de la secundaria donde curso'"`
	PromedioSecundaria     *float64 `gorm:"default:null;comment:'promedio obtenido en la secundaria'"`
	Tutor                  *string  `gorm:"default:null"`
	TutorAP                string   `gorm:"not null"`
	TutorAM                *string  `gorm:"default:null"`
	DireccionTutor         *string  `gorm:"default:null"`
	TelefonoTutor          *string  `gorm:"default:null"`
	EmailTutor             *string  `gorm:"default:null;comment:'email del tutor'"`
	Observaciones          *string  `gorm:"default:null;comment:'observaciones sobre el aspirante'"`
	Obserficha             string   `gorm:"not null"`
	// IDPficha               int       `gorm:"not null"`
	FechaFicha        string    `gorm:"not null"`
	PassFicha         string    `gorm:"not null"`
	CarFicha          int       `gorm:"not null"`
	Car2Ficha         int       `gorm:"not null"`
	Nivel             int       `gorm:"not null"`
	OpPor             string    `gorm:"not null"`
	Fecha             int       `gorm:"not null"`
	Extension         string    `gorm:"not null"`
	NumReferencia     int       `gorm:"not null"`
	Status            int       `gorm:"not null"`
	ObservacionesPago string    `gorm:"not null"`
	Acta              int       `gorm:"not null"`
	CURPP             int       `gorm:"not null"`
	Identi            int       `gorm:"not null"`
	Cer               int       `gorm:"not null"`
	ObsActa           string    `gorm:"not null"`
	ObsCer            string    `gorm:"not null"`
	ObsCURPP          string    `gorm:"not null"`
	ObsIdenti         string    `gorm:"not null"`
	Aspirante         Aspirante `gorm:"not null"`
}
