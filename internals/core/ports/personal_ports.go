package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IPersonalHandler interface {
	NuevoPersonal(c *fiber.Ctx) error
	TodosPersonal(c *fiber.Ctx) error
	UnPersonal(c *fiber.Ctx) error
	EditarPersonal(c *fiber.Ctx) error
	BorrarPersonal(c *fiber.Ctx) error
}

type IPersonalRepository interface {
	SavePersonal(personal *domains.Personal) error
	AllPersonal() ([]*domains.Personal, error)
	FindPersonalById(id int) (*domains.Personal, error)
	UpdatePersonal(personal *domains.Personal) error
	DeletePersonal(id int) error
}

type IPersonalService interface {
	GuardarPersonal(personal *domains.Personal) error
	ObtenerTodosPersonal() ([]*domains.Personal, error)
	ObtenerUnPersonal(id int) (*domains.Personal, error)
	ActualizarPersonal(personal *domains.Personal) error
	EliminarPersonal(id int) error
}
