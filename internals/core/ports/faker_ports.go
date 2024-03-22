package ports

import (
	"github.com/gofiber/fiber/v2"
)

type IFakerHandler interface {
	NuevoFaker(c *fiber.Ctx) error
	// TodosFakers(c *fiber.Ctx) error
	// UnFaker(c *fiber.Ctx) error
	// UnFakerGrupo(c *fiber.Ctx) error
	// UnFakerCalificaciones(c *fiber.Ctx) error
	// EditarFaker(c *fiber.Ctx) error
	// BorrarFaker(c *fiber.Ctx) error
}

type IFakerRepository interface {
	SaveFaker() error
	// AllFakers() ([]*domains.FakerSinNada, error)
	// FindFakerById(id int) (*domains.Faker, error)
	// FindFakerGrupoById(id int) (*domains.Faker, error)
	// FindFakerCalificacionesById(id int) (*domains.Faker, error)
	// UpdateFaker(faker *domains.Faker) error
	// DeleteFaker(id int) error
}

type IFakerService interface {
	GuardarFaker() error
	// ObtenerTodosFakers() ([]*domains.FakerSinNada, error)
	// ObtenerUnFaker(id int) (*domains.Faker, error)
	// ObtenerUnFakerGrupo(id int) (*domains.Faker, error)
	// ObtenerUnFakerCalificaciones(id int) (*domains.Faker, error)
	// ActualizarFaker(faker *domains.Faker) error
	// EliminarFaker(id int) error
}
