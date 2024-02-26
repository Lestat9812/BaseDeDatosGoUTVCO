package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IMateriaHandler interface {
	NuevaMateria(c *fiber.Ctx) error
	TodasMaterias(c *fiber.Ctx) error
	UnaMateria(c *fiber.Ctx) error
	EditarMateria(c *fiber.Ctx) error
	BorrarMateria(c *fiber.Ctx) error
}

type IMateriaRepository interface {
	SaveMateria(materia *domains.Materia) error
	AllMaterias() ([]*domains.Materia, error)
	FindMateriaById(id int) (*domains.Materia, error)
	UpdateMateria(materia *domains.Materia) error
	DeleteMateria(id int) error
}

type IMateriaService interface {
	GuardarMateria(materia *domains.Materia) error
	ObtenerTodasMaterias() ([]*domains.Materia, error)
	ObtenerUnaMateria(id int) (*domains.Materia, error)
	ActualizarMateria(materia *domains.Materia) error
	EliminarMateria(id int) error
}
