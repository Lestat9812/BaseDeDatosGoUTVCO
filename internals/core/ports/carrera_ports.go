package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type ICarreraHandler interface {
	NuevaCarrera(c *fiber.Ctx) error
	TodasCarreras(c *fiber.Ctx) error
	UnaCarrera(c *fiber.Ctx) error
	EditarCarrera(c *fiber.Ctx) error
	BorrarCarrera(c *fiber.Ctx) error
}

type ICarreraRepository interface {
	SaveCarrera(carrera *domains.Carrera) error
	AllCarreras() ([]*domains.Carrera, error)
	FindCarreraById(id int) (*domains.Carrera, error)
	UpdateCarrera(carrera *domains.Carrera) error
	DeleteCarrera(id int) error
}

type ICarreraService interface {
	GuardarCarrera(carrera *domains.Carrera) error
	ObtenerTodasCarreras() ([]*domains.Carrera, error)
	ObtenerUnaCarrera(id int) (*domains.Carrera, error)
	ActualizarCarrera(carrera *domains.Carrera) error
	EliminarCarrera(id int) error
}
