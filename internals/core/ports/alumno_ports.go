package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/adapters/repositories"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IAlumnoHandler interface {
	NuevoAlumno(c *fiber.Ctx) error
	TodosAlumnos(c *fiber.Ctx) error
	UnAlumno(c *fiber.Ctx) error
	EditarAlumno(c *fiber.Ctx) error
	BorrarAlumno(c *fiber.Ctx) error
}

type IAlumnoRepository interface {
	SaveAlumno(alumno *domains.Alumno) error
	AllAlumnos() ([]*domains.Alumno, error)
	FindAlumnoById(id int) (*repositories.AlumnoSinNada, error)
	UpdateAlumno(alumno *domains.Alumno) error
	DeleteAlumno(id int) error
}

type IAlumnoService interface {
	GuardarAlumno(alumno *domains.Alumno) error
	ObtenerTodosAlumnos() ([]*domains.Alumno, error)
	ObtenerUnAlumno(id int) (*repositories.AlumnoSinNada, error)
	ActualizarAlumno(alumno *domains.Alumno) error
	EliminarAlumno(id int) error
}
