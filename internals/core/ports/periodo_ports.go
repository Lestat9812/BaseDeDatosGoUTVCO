package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IPeriodoHandler interface {
	NuevoPeriodo(c *fiber.Ctx) error
	TodosPeriodos(c *fiber.Ctx) error
	UnPeriodo(c *fiber.Ctx) error
	EditarPeriodo(c *fiber.Ctx) error
	BorrarPeriodo(c *fiber.Ctx) error
}

type IPeriodoRepository interface {
	SavePeriodo(periodo *domains.Periodo) error
	AllPeriodos() ([]*domains.Periodo, error)
	FindPeriodoById(id int) (*domains.Periodo, error)
	UpdatePeriodo(periodo *domains.Periodo) error
	DeletePeriodo(id int) error
}

type IPeriodoService interface {
	GuardarPeriodo(periodo *domains.Periodo) error
	ObtenerTodosPeriodos() ([]*domains.Periodo, error)
	ObtenerUnPeriodo(id int) (*domains.Periodo, error)
	ActualizarPeriodo(periodo *domains.Periodo) error
	EliminarPeriodo(id int) error
}
