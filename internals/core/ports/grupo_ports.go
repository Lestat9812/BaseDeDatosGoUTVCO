package ports

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/gofiber/fiber/v2"
)

type IGrupoHandler interface {
	NuevoGrupo(c *fiber.Ctx) error
	TodosGrupos(c *fiber.Ctx) error
	UnGrupo(c *fiber.Ctx) error
	EditarGrupo(c *fiber.Ctx) error
	BorrarGrupo(c *fiber.Ctx) error
}

type IGrupoRepository interface {
	SaveGrupo(grupo *domains.Grupo) error
	AllGrupos() ([]*domains.Grupo, error)
	FindGrupoById(id int) (*domains.Grupo, error)
	UpdateGrupo(grupo *domains.Grupo) error
	DeleteGrupo(id int) error
}

type IGrupoService interface {
	GuardarGrupo(grupo *domains.Grupo) error
	ObtenerTodosGrupos() ([]*domains.Grupo, error)
	ObtenerUnGrupo(id int) (*domains.Grupo, error)
	ActualizarGrupo(grupo *domains.Grupo) error
	EliminarGrupo(id int) error
}
