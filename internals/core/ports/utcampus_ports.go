package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IUtcampusHandler interface {
	NuevoUtcampus(c *fiber.Ctx) error
	TodosUtcampus(c *fiber.Ctx) error
	UnUtcampus(c *fiber.Ctx) error
	EditarUtcampus(c *fiber.Ctx) error
	BorrarUtcampus(c *fiber.Ctx) error
}

type IUtcampusRepository interface {
	SaveUtcampus(utcampus *domains.Utcampus) error
	AllUtcampus() ([]*domains.Utcampus, error)
	FindUtcampusById(id int) (*domains.Utcampus, error)
	UpdateUtcampus(utcampus *domains.Utcampus) error
	DeleteUtcampus(id int) error
}

type IUtcampusService interface {
	GuardarUtcampus(utcampus *domains.Utcampus) error
	ObtenerTodosUtcampus() ([]*domains.Utcampus, error)
	ObtenerUnUtcampus(id int) (*domains.Utcampus, error)
	ActualizarUtcampus(utcampus *domains.Utcampus) error
	EliminarUtcampus(id int) error
}
