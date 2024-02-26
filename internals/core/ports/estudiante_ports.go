package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IEstudianteHandler interface {
	NuevoEstudiante(c *fiber.Ctx) error
	TodosEstudiantes(c *fiber.Ctx) error
	UnEstudiante(c *fiber.Ctx) error
	EditarEstudiante(c *fiber.Ctx) error
	BorrarEstudiante(c *fiber.Ctx) error
}

type IEstudianteRepository interface {
	SaveEstudiante(estudiante *domains.Estudiante) error
	AllEstudiantes() ([]*domains.Estudiante, error)
	FindEstudianteById(id int) (*domains.Estudiante, error)
	UpdateEstudiante(estudiante *domains.Estudiante) error
	DeleteEstudiante(id int) error
}

type IEstudianteService interface {
	GuardarEstudiante(estudiante *domains.Estudiante) error
	ObtenerTodosEstudiantes() ([]*domains.Estudiante, error)
	ObtenerUnEstudiante(id int) (*domains.Estudiante, error)
	ActualizarEstudiante(estudiante *domains.Estudiante) error
	EliminarEstudiante(id int) error
}
